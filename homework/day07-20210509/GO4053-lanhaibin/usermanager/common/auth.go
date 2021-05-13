package common

import (
	"crypto/sha512"
	"fmt"

	"github.com/princebot/getpass"
)

var PASSWORD = "ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413"

func Login() bool {
	password, err := getpass.Get("请输入密码:")
	if err != nil {
		return false
	}
	hash := sha512.New()
	hash.Write(password)
	return fmt.Sprintf("%x", hash.Sum([]byte{})) == PASSWORD
}
