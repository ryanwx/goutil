// All crypt methods of goutil return the []byte data type.
// Usually, we usually use the string type after base64 as the transfer.
// You only need to call the method base64.StdEncoding.EncodeToString after getting the cipherText.
// Similarly, if you need to decrypt the string after base64, you need to call base64.StdEncoding.DecodeString
// to get the cipherText of []byte data structure before calling the corresponding crypt.
package goutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"hash"
	"strings"
)

// cipher block struct
type cipherBlock struct {
	block cipher.Block
	size  int
}

// encrypt a []byte slice to a upper MD5 string.
func Md5(plain []byte) string {
	h := md5.New()

	return hexEncrypt(h, plain)
}

// encrypt a []byte slice to a upper sha1 string.
func Sha1(plain []byte) string {
	h := sha1.New()

	return hexEncrypt(h, plain)
}

// encrypt a []byte slice to a upper sha256 string.
func Sha256(plain []byte) string {
	h := sha256.New()

	return hexEncrypt(h, plain)
}

// AES CBC encrypt
// It will create a random iv.
// And iv append to the encrypted cipherText, returning the byte array after their combination.
func EncryptCBCAESJoinIv(cipherKey []byte, src []byte) ([]byte, error) {
	iv, err := randIv(len(cipherKey))
	if nil != err {
		return nil, err
	}

	dst, err := EncryptCBCAES(cipherKey, src, iv)
	if nil != err {
		return nil, err
	}

	return append(dst, iv...), nil
}

// AES CBC Decrypt.
// It can decrypt the cipherText combined with the byte array of iv.
func DecryptCBCAESJoinIv(cipherKey []byte, dst []byte) ([]byte, error) {
	ivStartInd := len(dst) - len(cipherKey)

	return DecryptCBCAES(cipherKey, dst[:ivStartInd], dst[ivStartInd:])
}

// AES CBC Encrypt with IV.
// You have to provide it with a clear iv.
// It returns the encrypted cipherText and does not mix any other data.
func EncryptCBCAES(cipherKey []byte, src []byte, iv []byte) ([]byte, error) {
	cb, err := newCipherBlock(cipherKey)
	if err != nil {
		return nil, err
	}
	src = aesPadding(src, cb)
	blockMode := cipher.NewCBCEncrypter(cb.block, iv)

	return aesBlockCrypt(blockMode, src), nil
}

// AES CBC Decrypt
// Must explicitly declare the decrypted iv.
// The decrypted object should be a pure encrypted cipherText, instead of adding any other data after the cipherText.
func DecryptCBCAES(cipherKey []byte, dst []byte, iv []byte) ([]byte, error) {
	cb, err := newCipherBlock(cipherKey)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(cb.block, iv)

	return aesUnPadding(aesBlockCrypt(blockMode, dst)), nil
}

// Rsa Pkcs1 Encrypt
func EncryptRsaPkcs1(plainText string, publicKey []byte) (string, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("failed to Decode PubKey")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", errors.New("failed to parse pubKey")
	}
	pub := pubInterface.(*rsa.PublicKey)
	cipherData, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(plainText))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherData), nil
}

// Rsa Pkcs1 Decrypt
func DecryptRsaPkcs1(cipherText string, privateKey []byte) (string, error) {
	cipherData, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("failed to decode private key")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	plainData, err := rsa.DecryptPKCS1v15(rand.Reader, priv, []byte(cipherData))
	if err != nil {
		return "", err
	}

	return string(plainData), nil
}

// write []byte to hash object, and checksum it.
func hexEncrypt(h hash.Hash, plain []byte) string {
	h.Write(plain)

	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// new cipher Block.
func newCipherBlock(cipherKey []byte) (cipherBlock, error) {
	block, err := aes.NewCipher(cipherKey)
	if nil != err {
		return cipherBlock{}, err
	}

	blockSize := block.BlockSize()

	return cipherBlock{block: block, size: blockSize}, nil
}

// rand a iv.
func randIv(length int) ([]byte, error) {
	iv := make([]byte, length)
	n, err := rand.Read(iv)
	if nil != err || n != length {
		return nil, err
	}

	return iv, nil
}

// aes crypt block, Encrypt or Decrypt.
func aesBlockCrypt(blockMode cipher.BlockMode, src []byte) []byte {
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)

	return dst
}

// aes cipherText padding func.
func aesPadding(cipherText []byte, cb cipherBlock) []byte {
	padding := cb.size - len(cipherText)%cb.size
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, paddingText...)
}

// aes cipherText unPadding func.
func aesUnPadding(originText []byte) []byte {
	length := len(originText)
	unPadding := int(originText[length-1])

	return originText[:(length - unPadding)]
}
