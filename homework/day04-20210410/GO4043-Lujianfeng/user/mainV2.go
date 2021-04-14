package main

import (
	"fmt"

	"user/user"

	"github.com/princebot/getpass"
)

const Pass = "123456"

func input(prompt string) string {
	var txt string
	fmt.Scan(&txt)
	return txt
}
func prompt() {
	fmt.Println(`
	"欢迎使用XXX系统"
	1. 添加用户 ==> add
	2. 删除用户 ==> del
	3. 修改用户 ==> modify
	4. 查询用户 ==> query
	5. 退出 ==> quit
	`)
}

func main() {

	//欢迎进入系统
	// var password string
	// fmt.Print("请输入密码：")
	// fmt.Scan(&password)
	// fmt.Println("你输入的密码是：", password)
	// fmt.Println(getpass.Get("请输入密码："))

	passwordByte, _ := getpass.Get("请输入密码:")
	fmt.Println("你输入的密码是：", string(passwordByte))

	if string(passwordByte) == Pass {
		fmt.Println("密码正确，欢迎进入XXX系统")
	} else {
		// for i := 0; i < 3; i++ {
		// 	fmt.Println("密码错误，请重新输入")
		// }
		fmt.Println("密码错误，请重新输入")

	}

END:
	for {
		// var choiceOne string
		// fmt.Printf("请选择你要做的操作: ")
		// fmt.Scan(&choiceOne)
		prompt()
		cmd := input("请输入指令：")
		switch cmd {
		case "1":
			user.Add()
		case "2":

			user.Del()
		case "3":

			user.Mod()
		case "4":

			user.Query()
		case "5":
			break END
		default:
			fmt.Println("指令错误")
		}
	}

}
