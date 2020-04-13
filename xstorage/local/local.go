package local

import (
	"errors"
	"io/ioutil"
	"os"
)

type StorageLocal struct {
	savePath string
}

var localObject *StorageLocal

func GetStorageLocal() (*StorageLocal, error) {
	if localObject == nil {
		err := LocalInit()
		if err != nil {
			return nil, err
		}
	}

	return localObject, nil
}

func LocalInit() error {
	return nil
}

func (s *StorageLocal) Init() error {
	return nil
}

func (s *StorageLocal) PutObject(filename string, value []byte) (string, error) {
	if len(value) == 0 {
		return "", errors.New("param[value] is empty")
	}

	if filename == "" {
		return "", errors.New("param[filename] is empty")
	}

	err := s.Init()
	if err != nil {
		return "", err
	}

	writeFile, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer writeFile.Close()

	_, err = writeFile.Write(value)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (s *StorageLocal) GetObject(filename string) ([]byte, error) {
	if filename == "" {
		return nil, errors.New("param[filename] is empty")
	}

	err := s.Init()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(filename)
}

func (s *StorageLocal) GetStoragePath(fileName string) string {
	return fileName
}
