package ziputil

import (
	"archive/zip"
	"bytes"
	"github.com/pubgo/x/xerror"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// ZipFile represents a zip file.
type ZipFile struct {
	zipFile *os.File
	writer  *zip.Writer
}

// Create creates a zip file with the specified filename.
func Create(filename string) (z *ZipFile, err error) {
	defer xerror.RespErr(&err)
	file := xerror.PanicErr(os.Create(filename)).(*os.File)
	return &ZipFile{zipFile: file, writer: zip.NewWriter(file)}, nil
}

// Close closes the zip file writer.
func (z *ZipFile) Close() (err error) {
	defer xerror.RespErr(&err)
	xerror.Panic(z.writer.Close())
	return z.zipFile.Close() // close the underlying writer
}

// AddEntryN adds entries.
func (z *ZipFile) AddEntryN(path string, names ...string) (err error) {
	defer xerror.RespErr(&err)

	for _, name := range names {
		xerror.Panic(z.AddEntry(filepath.Join(path, name), name))
	}

	return
}

// AddEntry adds a entry.
func (z *ZipFile) AddEntry(path, name string) (err error) {
	defer xerror.RespErr(&err)

	fi := xerror.PanicErr(os.Stat(name)).(os.FileInfo)
	fh := xerror.PanicErr(zip.FileInfoHeader(fi)).(*zip.FileHeader)
	fh.Name = filepath.ToSlash(filepath.Clean(path))
	fh.Method = zip.Deflate // data compression algorithm

	if fi.IsDir() {
		fh.Name = fh.Name + "/" // be care the ending separator
	}

	file := xerror.PanicErr(os.Open(name)).(*os.File)
	defer file.Close()
	xerror.PanicErr(io.Copy(xerror.PanicErr(z.writer.CreateHeader(fh)).(io.Writer), file))

	return
}

// AddDirectoryN adds directories.
func (z *ZipFile) AddDirectoryN(path string, names ...string) (err error) {
	defer xerror.RespErr(&err)

	for _, name := range names {
		xerror.Panic(z.AddDirectory(path, name))
	}

	return
}

// AddDirectory adds a directory.
func (z *ZipFile) AddDirectory(path, dirName string) error {
	files, err := ioutil.ReadDir(dirName)
	if nil != err {
		return err
	}

	if 0 == len(files) {
		err := z.AddEntry(path, dirName)
		if nil != err {
			return err
		}

		return nil
	}

	for _, file := range files {
		localPath := filepath.Join(dirName, file.Name())
		zipPath := filepath.Join(path, file.Name())

		err = nil
		if file.IsDir() {
			err = z.AddDirectory(zipPath, localPath)
		} else {
			err = z.AddEntry(zipPath, localPath)
		}

		if nil != err {
			return err
		}
	}

	return nil
}

func cloneZipItem(f *zip.File, dest string) error {
	// create full directory path
	fileName := f.Name

	if !utf8.ValidString(fileName) {
		data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(fileName)), simplifiedchinese.GB18030.NewDecoder()))
		if nil == err {
			fileName = string(data)
		}
	}

	path := filepath.Join(dest, fileName)

	err := os.MkdirAll(filepath.Dir(path), os.ModeDir|os.ModePerm)
	if nil != err {
		return err
	}

	if f.FileInfo().IsDir() {
		err = os.Mkdir(path, os.ModeDir|os.ModePerm)
		if nil != err {
			return err
		}

		return nil
	}

	// clone if item is a file

	rc, err := f.Open()
	if nil != err {
		return err
	}

	defer rc.Close()

	// use os.Create() since Zip don't store file permissions
	fileCopy, err := os.Create(path)
	if nil != err {
		return err
	}

	defer fileCopy.Close()

	_, err = io.Copy(fileCopy, rc)
	if nil != err {
		return err
	}

	return nil
}

// Unzip extracts a zip file specified by the zipFilePath to the destination.
func Unzip(zipFilePath, destination string) error {
	r, err := zip.OpenReader(zipFilePath)

	if nil != err {
		return err
	}

	defer r.Close()

	for _, f := range r.File {
		err = cloneZipItem(f, destination)
		if nil != err {
			return err
		}
	}

	return nil
}

//压缩文件
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址
func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压
func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}


func IsZip(zipPath string) bool {
	f, err := os.Open(zipPath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

func Unzip1(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}
		//------------注入

		dir := filepath.Dir(path)
		if len(dir) > 0 {
			if _, err = os.Stat(dir); os.IsNotExist(err) {
				err = os.MkdirAll(dir, 0755)
				if err != nil {
					return err
				}
			}
		}

		//---------------------end

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

//压缩文件
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址
func Compress1(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		if len(prefix) == 0 {
			prefix = info.Name()
		} else {
			prefix = prefix + "/" + info.Name()
		}
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		if len(prefix) == 0 {
			header.Name = header.Name
		} else {
			header.Name = prefix + "/" + header.Name
		}
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
