package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	uuidFormat = "%s-%s"
)

func Md5(data string) string {
	d := []byte(data)
	m := md5.New()
	m.Write(d)

	return hex.EncodeToString(m.Sum(nil))
}

func Sha256(data string) string {
	d := []byte(data)
	m := sha256.New()
	m.Write(d)

	return hex.EncodeToString(m.Sum(nil))
}

func SpanId() string {
	now := time.Now().UnixNano()
	rand.Seed(now)
	random := strconv.FormatInt(rand.Int63n(999999999), 10)
	return Md5(fmt.Sprintf(uuidFormat, random, strconv.FormatInt(now, 10)))
}

func TraceId() string {
	now := time.Now().UnixNano()
	rand.Seed(now)
	random := strconv.FormatInt(rand.Int63n(999999999), 10)
	return Sha256(fmt.Sprintf(uuidFormat, random, strconv.FormatInt(now, 10)))
}

func IsDir(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return false
	}

	return p.IsDir()
}

func IsFile(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !p.IsDir()
}
