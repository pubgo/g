package pathutil

import (
	"errors"
	"os"

	"path/filepath"
)

// Exist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func Exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

// PathExists path is exist
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Search Search a file in paths.
func Search(fileName string, paths ...string) (fullpath string, err error) {
	for _, path := range paths {
		if fullpath = filepath.Join(path, fileName); Exist(fullpath) {
			return
		}
	}
	err = errors.New(fullpath + " not found in paths")
	return
}
