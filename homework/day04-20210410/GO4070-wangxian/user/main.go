package main

import (
	"fmt"
	"strings"
	"user/command"
	_ "user/init"
	"user/utils"

	"github.com/princebot/getpass"
)

var commands = command.Commands

const (
	user     = "admin"
	password = "a3dcb8aae0d84122a1778e26da285857" //123456
)

//用户登录
func Login() bool {
	for i := 0; i < 3; i++ {
		name := utils.Input("请输入用户名:")
		pw, _ := getpass.Get("请输入密码:")

		if name == "admin" && utils.Md5Encrypt(pw) == password {
			fmt.Println("登录成功!\n")
			return true
		} else {
			fmt.Printf("用户名或密码错误, 还可尝试%d次\n\n", 2-i)
		}
	}
	return false
}

//打印提示信息
func prompt() {
	propmtInfo := command.Prompt

	fmt.Println(strings.Repeat("*", 15))
	fmt.Println("欢迎使用用户管理系统")
	fmt.Println("1 退出")
	for _, v := range propmtInfo {
		fmt.Printf("%s %s\n", v[0], v[1])
	}
	fmt.Println(strings.Repeat("*", 15))
}

func main() {
	if Login() {
	END:
		for {
			prompt()
			cmd := utils.Input("请输入指令:")

			if command, ok := commands[cmd]; ok {
				command()
			} else if cmd == "1" {
				break END
			} else {
				fmt.Println("指令错误")
			}
		}
	} else {
		fmt.Println("登录失败，请稍后再试!")
	}

}
