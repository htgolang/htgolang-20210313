package cred

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Verify(passwd, secrect string) bool {
	h := md5.New()
	io.WriteString(h, passwd)

	if fmt.Sprintf("%x", h.Sum(nil)) == secrect {
		return true
	} else {
		return false
	}

}
