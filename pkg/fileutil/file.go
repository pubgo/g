// doc https://github.com/spf13/afero

package fileutil

import (
	"bufio"
	"github.com/pubgo/g/pkg/encoding/hashutil"
	"github.com/pubgo/g/xerror"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// GetSize get the file size
func GetSize(r io.Reader) (i int, err error) {
	defer xerror.RespErr(&err)

	content, err := ioutil.ReadAll(r)
	xerror.Panic(err)

	i = len(content)
	return
}

// GetExt get the file ext
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckNotExist check if the file exists
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// CheckPermission check if the file has permission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func IsNotExistMkDir(src string) (err error) {
	defer xerror.RespErr(&err)

	if notExist := CheckNotExist(src); notExist == true {
		xerror.PanicM(MkDir(src), "MkDir Error")
	}

	return
}

// MkDir create a directory
func MkDir(src string) (err error) {
	return xerror.Wrap(os.MkdirAll(src, os.ModePerm), "os.MkdirAll Error")
}

// Open a file according to a specific mode
func Open(name string, flag int, perm os.FileMode) (f *os.File, err error) {
	defer xerror.RespErr(&err)
	f = xerror.PanicErr(os.OpenFile(name, flag, perm)).(*os.File)
	return
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (f *os.File, err error) {
	defer xerror.RespErr(&err)

	dir, err := os.Getwd()
	xerror.PanicM(err, "os.Getwd Error")

	src := filepath.Join(dir, filePath)
	perm := CheckPermission(src)
	xerror.PanicT(perm, "file.CheckPermission Permission denied src: %s", src)
	xerror.PanicM(IsNotExistMkDir(src), "file.IsNotExistMkDir src: %s", src)

	f = xerror.PanicErr(Open(filepath.Join(src, fileName), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)).(*os.File)
	return
}

// GetImageName get image hash name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = hashutil.MD5(fileName)
	return fileName + ext
}

func FileExt(filePathName string) string {
	return strings.TrimPrefix(filepath.Ext(filePathName), ".")
}

// GetFileSize get the length in bytes of file of the specified path.
func GetFileSize(path string) (i int64, err error) {
	defer xerror.RespErr(&err)
	fi := xerror.PanicErr(os.Stat(path)).(os.FileInfo)
	i = fi.Size()
	return
}

// IsExist determines whether the file spcified by the given path is exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsBinary determines whether the specified content is a binary file content.
func IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}
	return false
}

// IsImg determines whether the specified extension is a image.
func IsImg(extension string) bool {
	ext := strings.ToLower(extension)
	switch ext {
	case ".jpg", ".jpeg", ".bmp", ".gif", ".png", ".svg", ".ico":
		return true
	default:
		return false
	}
}

// IsDir determines whether the specified path is a directory.
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}

	if nil != err {
		return false
	}

	return fio.IsDir()
}

// CopyFile copies the source file to the dest file.
func CopyFile(source string, dest string) (err error) {
	defer xerror.RespErr(&err)

	sourcefile := xerror.PanicErr(os.Open(source)).(*os.File)
	defer sourcefile.Close()

	destfile := xerror.PanicErr(os.Create(dest)).(*os.File)
	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	xerror.Panic(err)

	sourceinfo := xerror.PanicErr(os.Stat(source)).(os.FileInfo)
	xerror.Panic(os.Chmod(dest, sourceinfo.Mode()))

	return
}

// CopyDir copies the source directory to the dest directory.
func CopyDir(source string, dest string) (err error) {
	defer xerror.RespErr(&err)

	sourceinfo, err := os.Stat(source)
	xerror.Panic(err)

	// create dest dir
	xerror.Panic(os.MkdirAll(dest, sourceinfo.Mode()))

	directory := xerror.PanicErr(os.Open(source)).(*os.File)
	defer directory.Close()

	objects := xerror.PanicErr(directory.Readdir(-1)).([]os.FileInfo)
	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			xerror.Panic(CopyDir(srcFilePath, destFilePath))
		} else {
			xerror.Panic(CopyFile(srcFilePath, destFilePath))
		}
	}

	return
}

// GrepFile like command grep -E
// for example: GrepFile(`^hello`, "hello.txt")
// \n is striped while read
func GrepFile(patten string, filename string) (lines []string, err error) {
	re, err := regexp.Compile(patten)
	if err != nil {
		return
	}

	fd, err := os.Open(filename)
	if err != nil {
		return
	}
	lines = make([]string, 0)
	reader := bufio.NewReader(fd)
	prefix := ""
	var isLongLine bool
	for {
		byteLine, isPrefix, er := reader.ReadLine()
		if er != nil && er != io.EOF {
			return nil, er
		}
		if er == io.EOF {
			break
		}
		line := string(byteLine)
		if isPrefix {
			prefix += line
			continue
		} else {
			isLongLine = true
		}

		line = prefix + line
		if isLongLine {
			prefix = ""
		}
		if re.MatchString(line) {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

// CheckFileIsExist 检查目录是否存在
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// BuildDir 创建目录
func BuildDir(absDir string) error {
	return os.MkdirAll(path.Dir(absDir), os.ModePerm) //生成多级目录
}

// DeleteFile 删除文件或文件夹
func DeleteFile(absDir string) error {
	return os.RemoveAll(absDir)
}

// ReadFile 读取文件
func ReadLine(path string) (src []string) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return []string{}
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		src = append(src, string(line))
	}

	return src
}

func DirName(argv ...string) string {
	file := ""
	if len(argv) > 0 && argv[0] != "" {
		file = argv[0]
	} else {
		file, _ = exec.LookPath(os.Args[0])
	}
	path, _ := filepath.Abs(file)
	directory := filepath.Dir(path)
	return strings.Replace(directory, "\\", "/", -1)
}

func GetProPath() string {
	return DirName("root")
}

func IsFile(f string) bool {
	if fi, err := os.Stat(f); err != nil {
		return false
	} else {
		return !fi.IsDir()
	}
}

func AppendFile(filename string, data string) {
	if !IsFile(filename) {
		//WriteFile(filename, data)
		return
	}
	if f, err := os.OpenFile(filename, os.O_WRONLY, 0644); err != nil {
		xerror.PanicM(err, "os.OpenFile错误")
	} else {
		if n, err := f.Seek(0, 2); err != nil {
			xerror.PanicM(err, "f.Seek错误")
		} else {
			if _, err = f.WriteAt([]byte(data), n); err != nil {
				xerror.PanicM(err, "f.WriteAt错误")
			}
		}
		_ = f.Close()
	}
}

func WriteString(path string, content string, append bool) error {
	flag := os.O_RDWR | os.O_CREATE
	if append {
		flag = flag | os.O_APPEND
	}
	file, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}

func AppendLine(path string, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	content = strings.Join([]string{content, "\n"}, "")
	_, err = file.WriteString(content)
	return err
}

// Read read file and return string
func ReadFile(path string) (string, error) {
	fin, err := os.Open(path)
	if err != nil {
		log.Println("os.Open: ", path, err)
		return "", err
	}
	defer fin.Close()

	var str string
	buf := make([]byte, 1024)
	for {
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
		// os.Stdout.Write(buf[:n])
		strBuf := string(buf[:n])
		str += strBuf
	}

	return str, nil
}

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it
// and its upper level paths.
func WriteFile(fileName string, data []byte) error {
	os.MkdirAll(path.Dir(fileName), os.ModePerm)
	return ioutil.WriteFile(fileName, data, 0655)
}

// Write writes data to a file named by filename.
// If the file does not exist, WriteFile creates it
// and its upper level paths.
func Write(fileName, writeStr string) {
	os.MkdirAll(path.Dir(fileName), os.ModePerm)

	fout, err := os.Create(fileName)
	if err != nil {
		log.Println("Write file "+fileName, err)
		return
	}
	defer fout.Close()

	fout.WriteString(writeStr)
}

// AppendTo append to file
func AppendTo(fileName, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		log.Println("File open failed. err: " + err.Error())
		return err
	}

	n, _ := f.Seek(0, os.SEEK_END)
	_, err = f.WriteAt([]byte(content), n)

	f.Close()
	return err
}

// Empty empty the file
func Empty(fileName string, args ...int64) {
	var size int64
	if len(args) > 0 {
		size = args[0]
	}

	os.Truncate(fileName, size)
}
