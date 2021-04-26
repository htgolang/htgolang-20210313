package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/princebot/getpass"
	"log"
)

const (
	MaxAuthCount = 3
	Password = "202cb962ac59075b964b07152d234b70"
)

func Auth() bool {
	for i := 0; i < MaxAuthCount; i++ {
		pwd, err := getpass.Get("请输入密码: ")
		if err != nil {
			log.Fatal(err)
		}
		if fmt.Sprintf("%x", md5.Sum(pwd)) == Password {
			return true
		} else {
			if i == MaxAuthCount-1 {
				fmt.Println("密码错误，您的三次机会已用完")
			} else {
				fmt.Printf("密码错误，您还有%d次机会\n", MaxAuthCount-i-1)
			}
		}
	}
	return false
}
