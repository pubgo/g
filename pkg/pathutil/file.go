package pathutil

import (
	"bufio"
	"github.com/pubgo/g/xerror"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

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

// IsExist determines whether the file spcified by the given path is exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
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

// GetPathDirs 获取目录所有文件夹
func GetPathDirs(absDir string) (re []string) {
	if CheckFileIsExist(absDir) {
		files, _ := ioutil.ReadDir(absDir)
		for _, f := range files {
			if f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// GetPathFiles 获取目录所有文件
func GetPathFiles(absDir string) (re []string) {
	if CheckFileIsExist(absDir) {
		files, _ := ioutil.ReadDir(absDir)
		for _, f := range files {
			if !f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// GetModelPath 获取目录地址
func GetModelPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path := filepath.Dir(file)
	path, _ = filepath.Abs(path)

	return path
}

// GetCurrentDirectory 获取程序运行路径
func GetCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
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

// List list file
func List(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, nil
}

// ListDir list dir
func ListDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, nil
}

// Walk walk file
func Walk(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)
	err = filepath.Walk(
		dirPth, func(filename string, fi os.FileInfo, err error) error {
			if fi.IsDir() {
				return nil
			}

			if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
				files = append(files, filename)
			}
			return nil
		})

	return files, err
}

// WalkDir walk dir
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)
	err = filepath.Walk(
		dirPth, func(filename string, fi os.FileInfo, err error) error {
			if !fi.IsDir() {
				return nil
			}

			if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
				files = append(files, filename)
			}
			return nil
		})

	return files, err
}

// Copy copies file from source to target path.
func Copy(src, dst string) error {
	// Gather file information to set back later.
	si, err := os.Lstat(src)
	if err != nil {
		return err
	}

	// Handle symbolic link.
	if si.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(src)
		if err != nil {
			return err
		}
		// NOTE: os.Chmod and os.Chtimes don't recoganize symbolic link,
		// which will lead "no such file or directory" error.
		return os.Symlink(target, dst)
	}

	sr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sr.Close()

	dw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dw.Close()

	if _, err = io.Copy(dw, sr); err != nil {
		return err
	}

	// Set back file information.
	if err = os.Chtimes(dst, si.ModTime(), si.ModTime()); err != nil {
		return err
	}
	return os.Chmod(dst, si.Mode())
}
