package main

import (
	"fmt"
	"strconv"
	"strings"

	"manager/commands"

	_ "manager/init"
	"manager/utils"

	"github.com/princebot/getpass"
)

const Password = "123" // md5

func prompt() {

	fmt.Println(strings.Repeat("*", 10))
	fmt.Println("欢迎使用xxx系统")
	fmt.Println("1. 退出")
	commands.Commands.Prompt(2, "%d. %s\n")
	fmt.Println(strings.Repeat("*", 10))
}

func main() {
	// 功能串起来

	// 启动后给欢迎
	// 输入密码 3次输入密码失败 退出

	// 输入密码
	passwordBytes, _ := getpass.Get("请输入密码:")

	if Password != string(passwordBytes) {
		fmt.Println("密码错误")
		return
	}

END:
	for {
		prompt()
		cmd, _ := strconv.Atoi(utils.Input("请输入指令:"))
		if command := commands.Commands.Get(cmd); command != nil {
			command()
		} else if cmd == 1 {
			break END
		} else {
			fmt.Println("指令错误")
		}
	}
}
