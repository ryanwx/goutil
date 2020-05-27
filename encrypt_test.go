package goutil

import (
	"encoding/base64"
	"errors"
	"fmt"
	"testing"
)

// 2048 bit test pkcs1 public key.
var testRsaPkcs1PubKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCgdqnRT/ZUk2BUcH3xGkIPn9nb
7Eh55aMT8eOIQreUmHOfPKr+iHRf3e5X6Lmkk+I9dmt1k/9hJ4fo5VyCP6OnJWdd
HkVj6GA1ozBwQChBBL2p258RUNomFEWmxccS9Fd5A6xufPk2zu2ddTTNPVJCDCnJ
YikjCQua4Go5X28QOQIDAQAB
-----END PUBLIC KEY-----
`)

// 2048 bit test pkcs1 private key.
var testRsaPkcs1PriKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCgdqnRT/ZUk2BUcH3xGkIPn9nb7Eh55aMT8eOIQreUmHOfPKr+
iHRf3e5X6Lmkk+I9dmt1k/9hJ4fo5VyCP6OnJWddHkVj6GA1ozBwQChBBL2p258R
UNomFEWmxccS9Fd5A6xufPk2zu2ddTTNPVJCDCnJYikjCQua4Go5X28QOQIDAQAB
AoGADKJzShod9vhYuiVWYUzUPUc0n/qplUGF1uzxobnBGzfqyLO98gc8BI1ktjLK
pUEIrBa+LNafod57NtUSdOX+ZGuXUMf0Pk/I5stKlICzBDYAX1JfT/1Dp03++ycz
FsYC+3W2X2DQyOYjtjfIaXOVB07JD+GOYkTYrGFqGkFBzoECQQDWeti67fLKrqbB
cLBwqaF6JZOaRdwaSFH88co2YjaawdG+r8dySygRY+uoNgXqKAvb4LIXTF0mp8n4
i3pfRITBAkEAv4bijyu45neSkXYjY4R7TOZhWkiCMRF4uhV1Cjdh0hZEy99XRZvH
LoTm3w8pwFuEXuzlGyY0V4v5+nyyKJuReQJAQZPSDGWQpJP9/Is+B1R2QOiYVsUh
ZQ/TsakkELi6xVqZjNol+zYrjBWnCglqiYuxBIuRKDp7CMSopkvPIK3MgQJAe+Au
IRPSX53u+o5CjVdeuHo5dT94lWwLfa/rJ1RyvIMStBocRDVhOsFS4erYwkVu9Eac
WFb5e7ZZVJ3aTVFxYQJAR7wvQCLHocWs1lKWjnzsumJUTArr9U3F8PXWNCRYEurb
yAgWicVkGlSQbRdEHojgBq/z0UEB9v43ahvD0w+uNQ==
-----END RSA PRIVATE KEY-----
`)

func TestMd5(t *testing.T) {
	fmt.Println(Md5([]byte("soiouwsnsljsA")))
}

func TestSha1(t *testing.T) {
	fmt.Println(Sha1([]byte("soiouwsnsljsA")))
}

func TestSha256(t *testing.T) {
	fmt.Println(Sha256([]byte("soiouwsnsljsA")))
}

func TestCBCCrypt(t *testing.T) {
	key := []byte("1234567890asdfgh")
	src := []byte("dgs=292")
	dst, err := EncryptCBCAESJoinIv(key, src)
	if nil != err {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(dst))

	src, err = DecryptCBCAESJoinIv(key, dst)
	if nil != err {
		panic(err)
	}

	fmt.Println(string(src))

}

func TestEncryptRsaPkcs1(t *testing.T) {
	plain := "testRsaEncryptPkcs1"
	cipherText, err := EncryptRsaPkcs1(plain, testRsaPkcs1PubKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(cipherText)
}

func TestDecryptRsaPkcs1(t *testing.T) {
	cipherText := "H5vVoQIHYWX1Agl6KZG25CNGZqeO8mom3+8umrzi4EidDQ4jOeSzXoXIFY6lZVNcSl8W89rLIeZfm9bduZyyIysEStfqAvm6uFnJu68b+LqXvxYz7BXBjc2IdedDVs7+72Cmjuf7M0tZQiUqrszf18AgP5dfez4HMJthMdhCIw4="
	plainText, err := DecryptRsaPkcs1(cipherText, testRsaPkcs1PriKey)
	if err != nil {
		panic(err)
	}

	if plainText != "testRsaEncryptPkcs1" {
		panic(errors.New("cipherText is wrong"))
	}

	fmt.Println(plainText)
}
