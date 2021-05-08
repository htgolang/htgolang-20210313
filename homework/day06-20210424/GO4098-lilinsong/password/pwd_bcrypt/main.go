package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 加密密码
func HashAndSalt(pwdStr string) (pwdHash string, err error) {
	pwd := []byte(pwdStr)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return
	}
	pwdHash = string(hash)
	return
}

// 验证密码
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}

func main() {

	passwordOK := "admin"
	passwordERR := "adminxx"

	hashStr, err := HashAndSalt(passwordOK)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hashStr)

	// 正确密码验证
	check := ComparePasswords(hashStr, passwordOK)
	if !check {
		fmt.Println("pw wrong")
	} else {
		fmt.Println("pw ok")
	}

	// 错误密码验证
	check = ComparePasswords(hashStr, passwordERR)
	if !check {
		fmt.Println("pw wrong")
	} else {
		fmt.Println("pw ok")
	}
}