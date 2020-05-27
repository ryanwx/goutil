package goutil

import (
	"fmt"
	"math/rand"
	"regexp"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestRandomString(t *testing.T) {
	l := 10
	randBytePool := []byte("+-.*/\\<>|()&^%$#@!QWERTYUIOPASDFGHJKLZXCVBNM")
	s := RandomString(randBytePool, l)
	regStr := fmt.Sprintf("[%s]{1,%d}", regexp.QuoteMeta(string(randBytePool)), l)
	t.Logf("regStr: %s", regStr)
	t.Logf("testing string: %s", s)
	reg, err := regexp.Compile(regStr)
	if err != nil {
		t.Errorf("failed to compile regexp, %v", err)
		t.FailNow()
	}

	if !reg.MatchString(s) {
		t.Errorf("failed to match string, %s", s)
		t.FailNow()
	}
}

func TestRandomStringEnAndInt(t *testing.T) {
	l := 20
	s := RandomStringEnAndInt(l)
	regStr := fmt.Sprintf("[%s]{1,%d}", randomPoolEnAndInt, l)
	t.Logf("regStr: %s", regStr)
	t.Logf("testing string : %s", s)
	reg, err := regexp.Compile(regStr)
	if err != nil {
		t.Errorf("failed to compile regexp, %v", err)
		t.FailNow()
	}

	if !reg.MatchString(s) {
		t.Errorf("failed to match string, %s", s)
		t.FailNow()
	}
}
