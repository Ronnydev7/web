package lib

import (
	"crypto/hmac"
	"crypto/sha256"
)

func GetHmac(secret []byte, data []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write(data)
	return h.Sum(nil)
}
