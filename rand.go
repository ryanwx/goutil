package goutil

import (
	"math/rand"
)

var randomPoolEnAndInt = []byte("0123456789QWERTYUIOPASDFGHJKLZXCVBNM")

// s
func RandomString(randomPool []byte, l int) string {
	var res []byte
	randomPoolLen := len(randomPool)
	for i := 0; i < l; i++ {
		res = append(res, randomPool[rand.Intn(randomPoolLen)])
	}

	return string(res)
}

func RandomStringEnAndInt(l int) string {
	return RandomString(randomPoolEnAndInt, l)
}
