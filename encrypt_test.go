package goutil

import (
    "encoding/base64"
    "fmt"
    "testing"
)

func TestMd5( t *testing.T){
    fmt.Println(Md5([]byte("soiouwsnsljsA")))
}

func TestSha1(t *testing.T){
    fmt.Println(Sha1([]byte("soiouwsnsljsA")))
}

func TestSha256(t *testing.T){
    fmt.Println(Sha256([]byte("soiouwsnsljsA")))
}

func TestCBCCrypt(t *testing.T) {
    key := []byte("1234567890asdfgh")
    src := []byte("dgs=292")
    dst, err := EncryptCBCAESJoinIv(key, src)
    if nil !=  err {
        panic(err)
    }
    fmt.Println(base64.StdEncoding.EncodeToString(dst))

    src, err = DecryptCBCAESJoinIv(key, dst)
    if nil != err {
        panic(err)
    }

    fmt.Println(string(src))

}
