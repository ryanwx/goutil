package goutil

import (
	"os"
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
