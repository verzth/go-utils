package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func GetTimebaseUnix(seconds int64) int64 {
	t := time.Now().Unix()
	return t - (t % seconds)
}

func GenerateSignature(key string, second ...int64) (signature string) {
	mac := hmac.New(sha256.New, []byte(key))
	var t int64
	if len(second) > 0 {
		t = GetTimebaseUnix(second[0])
	} else {
		t = GetTimebaseUnix(30)
	}
	mac.Write([]byte(fmt.Sprintf("%d", t)))
	sign := mac.Sum(nil)
	signature = hex.EncodeToString(sign)
	return
}

func VerifySignature(key string, signature string, second ...int64) bool {
	var t, sec int64
	if len(second) > 0 {
		sec = second[0]
		t = GetTimebaseUnix(sec)
	} else {
		sec = 30
		t = GetTimebaseUnix(sec)
	}
	sign := GenerateSignature(key, t)
	if sign == signature {
		return true
	} else {
		return compatibilityBackward(key, signature, t, sec)
	}
}

func compatibilityBackward(key string, signature string, t int64, sec int64) bool {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(fmt.Sprintf("%d", t-sec)))
	sign := mac.Sum(nil)
	return hex.EncodeToString(sign) == signature
}
