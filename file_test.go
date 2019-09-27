package goutil

import (
    "fmt"
    "testing"
)

func TestPathExists( t *testing.T){
    path := "./file.go"
    fmt.Println(PathExists(path))
}