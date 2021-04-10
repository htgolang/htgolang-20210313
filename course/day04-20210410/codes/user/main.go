package main

import (
	"fmt"
	"strings"

	"user/commands"

	_ "user/init"

	"github.com/princebot/getpass"
)

const Password = "123" // md5

func input(prompt string) string {
	fmt.Print(prompt)
	var txt string
	fmt.Scan(&txt)
	return txt
}

func prompt() {
	prompts := commands.Prompts

	fmt.Println(strings.Repeat("*", 10))
	fmt.Println("欢迎使用xxx系统")
	fmt.Println("1. 退出")
	for _, v := range prompts {
		fmt.Printf("%s. %s\n", v[0], v[1])
	}

	fmt.Println(strings.Repeat("*", 10))
}

func main() {
	// 功能串起来

	// 启动后给欢迎
	// 输入密码 3次输入密码失败 退出

	// 输入密码
	var password string
	fmt.Print("请输入密码:")
	fmt.Scan(&password)
	fmt.Println("你输入的密码:", password)

	passwordBytes, _ := getpass.Get("请输入密码:")

	fmt.Println("你输入的密码:", string(passwordBytes))
	// passwordBytes => md5 => 比较password

	// 循环
	// 提示用户输入指令 => 1. 添加， 2: 编辑, 3: 删除, 4: 查询, 5. 退出

	// commands := map[string]func(){
	// 	"1": user.Add,
	// 	"2": user.Modify,
	// 	"3": user.Delete,
	// 	"4": user.Query,
	// }

	commands := commands.Commands

END:
	for {
		prompt()
		cmd := input("请输入指令:")
		if command, ok := commands[cmd]; ok {
			command()
		} else if cmd == "1" {
			break END
		} else {
			fmt.Println("指令错误")
		}
	}
}
