package goutil

import (
	"os"
	"testing"
)

func TestPathExists(t *testing.T) {
	path := "./file.go"
	res, err := PathExists(path)
	if nil != err {
		t.Errorf("failed to check path exists, err[%v]", err)
		t.FailNow()
	}

	if !res {
		t.Error("test path not exists")
		t.FailNow()
	}
}

func TestCreateFileWithDir(t *testing.T) {
	filepath := "test/file/hello.log"
	dirPerm := os.FileMode(0422)
	f, err := CreateFileWithDir(filepath, dirPerm)
	if err != nil {
		t.Errorf("failed to create file, err[%v]", err)
		t.FailNow()
	}
	f.Close()
}
