package main

import (
	"fmt"
	"strconv"
	"strings"

	"mgr/commands"
	"mgr/utils"

	_ "mgr/init"

	"github.com/princebot/getpass"
)

const Password = "123" // md5

func prompt() {

	fmt.Println(strings.Repeat("*", 10))
	fmt.Println("欢迎使用xxx系统")
	fmt.Println("1. 退出")
	commands.Commands.Prompts("%d. %s\n")
	fmt.Println(strings.Repeat("*", 10))
}

func main() {
	// 功能串起来

	// 启动后给欢迎
	// 输入密码 3次输入密码失败 退出

	// 输入密码
	passwordBytes, _ := getpass.Get("请输入密码:")

	fmt.Println("你输入的密码:", string(passwordBytes))

END:
	for {
		prompt()
		cmd, err := strconv.Atoi(utils.Input("请输入指令:"))
		if err == nil {
			command := commands.Commands.Get(cmd)
			if command != nil {
				command()
			} else if cmd == 1 {
				break END
			} else {
				fmt.Println("指令错误")
			}
		} else {

			fmt.Println("指令错误")
		}

	}
}
