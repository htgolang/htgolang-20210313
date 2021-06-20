package utils

import "encoding/base64"

func Base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Base64Decode(t string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(t)
}
