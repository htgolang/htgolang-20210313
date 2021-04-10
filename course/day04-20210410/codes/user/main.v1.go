package main

import (
	"fmt"

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
	fmt.Println(`
欢迎使用xxx管理系统
1. 添加
2. xxx
3. xxx
4. xxx
5. xxx
	`)
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

END:
	for {
		prompt()
		cmd := input("请输入指令:")
		switch cmd {
		case "1":
			fmt.Println("add")
		case "2":
			fmt.Println("edit")
		case "3":
			fmt.Println("delete")
		case "4":
			fmt.Println("query")
		case "5":
			break END
		default:
			fmt.Println("指令错误")
		}
	}
}
