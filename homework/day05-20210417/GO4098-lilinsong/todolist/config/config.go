package config

import (
	"crypto/sha1"
	"fmt"
	"github.com/creachadair/getpass"
	"todolist/commands"
	"todolist/utils/ioutlis"
)

var candyPassword string = "7c4a8d09ca3762af61e59520943dc26494f8941b"

var LoginRetry int = 3

const TimeLayout = "2006-01-02 15:04:05"

func LoginAuth() bool {
	for i := LoginRetry; i > 0; i-- {
		name := ioutlis.Input("请输入用户名称: ")
		pass, _ := getpass.Prompt("请输入密码")
		hashPass := hashText(pass)
		user, flag, err := commands.GetUser(name)
		if err != nil {
			fmt.Println(err)
			return false
		}
		if flag {
			if user == "candy" && hashPass == candyPassword {
				return true
			}

		} else {
			if user == "tank" && pass == "123456" {
				return true
			}
		}
		if i != 1 {
			ioutlis.Error(fmt.Sprintf("输入的密码错误，剩余登录%d次数", i-1))
		}
	}
	ioutlis.Error(fmt.Sprintf("密码错误唱过%d次，程序退出", LoginRetry))
	return false
}

func hashText(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
