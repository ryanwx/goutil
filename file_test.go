package goutil

import (
	"testing"
)

func TestPathExists(t *testing.T) {
	path := "./file.go"
	res, err := PathExists(path)
	if nil != err {
		t.Errorf("failed to check path exists, %v", err)
		t.FailNow()
	}

	if !res {
		t.Error("test path not exists")
		t.FailNow()
	}
}
