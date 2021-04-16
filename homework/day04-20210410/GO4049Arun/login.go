package main

import (
	"crypto/md5"
	"fmt"

	"github.com/princebot/getpass"
)

func login() bool{
	passwordBytes, _ := getpass.Get("登录密码:")
	md5Pwd := fmt.Sprintf("%x", md5.Sum([]byte(passwordBytes)))

	str := "123"
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	if md5str == md5Pwd {
		fmt.Println("恭喜您登录成功,登录密码为:", string(passwordBytes))
		return true
	}else {
		return false
	}
}
