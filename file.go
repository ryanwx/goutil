package goutil

import (
	"os"
	"path"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if nil == err {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func CreateFileWithDir(filepath string, dirPerm os.FileMode) (*os.File, error) {
	dir := path.Dir(filepath)
	ok, err := PathExists(dir)
	if err != nil {
		return nil, err
	}
	if !ok {
		err = os.MkdirAll(dir, dirPerm)
		if err != nil {
			return nil, err
		}
	}
	return os.Create(filepath)
}
