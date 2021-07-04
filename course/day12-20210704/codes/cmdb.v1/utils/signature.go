package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

var (
	key = []byte("65e9ee4f2da48eabbc24e5b3ba7cc616")
)

func Signature(b []byte) string {
	hasher := hmac.New(sha256.New, key)
	return fmt.Sprintf("%x", hasher.Sum(b))
}
