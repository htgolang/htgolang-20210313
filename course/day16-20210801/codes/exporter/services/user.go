package services

import (
	"crypto/md5"
	"fmt"
	"strings"
)

var users = map[string]string{
	"kk":    "E10ADC3949BA59ABBE56E057F20F883E",
	"admin": "E10ADC3949BA59ABBE56E057F20F883E",
}

func ValidateUser(authorization string) bool {
	pos := strings.Index(authorization, ":")
	if pos < 0 {
		return false
	}
	user, password := authorization[:pos], authorization[pos+1:]
	if pwd, ok := users[user]; ok {
		// pwd // MD5(配置/数据)
		return pwd == fmt.Sprintf("%X", md5.Sum([]byte(password)))
	}
	return false
}
