package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charsetUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charsetLower = "abcdefghijklmnopqrstuvwxyz0123456789"
const charsetHexLower = "0123456789abcdef"
const charsetHexUpper = "0123456789ABCDEF"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset)-1)]
	}
	return string(b)
}

func RandString(length int) string {
	return RandStringWithCharset(length, charset)
}

func RandStringUpper(length int) string {
	return RandStringWithCharset(length, charsetUpper)
}

func RandStringLower(length int) string {
	return RandStringWithCharset(length, charsetLower)
}

func RandomHex(n int) string {
	return RandStringWithCharset(n, charsetHexLower)
}

func RandomHexUpper(n int) string {
	return RandStringWithCharset(n, charsetHexUpper)
}
