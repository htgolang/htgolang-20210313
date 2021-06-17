package main

import (
	"crypto/md5"
	"fmt"
	"github.com/howeyc/gopass"
	"os"
	"strings"
	"user/commands"
	_ "user/init"
)

const pass = "123"

func prompt() {
	fmt.Println(strings.Repeat("*", 20))
	fmt.Println("欢迎使用用户管理系统")
	fmt.Println("1.退出")
	prompts := commands.Prompt
	for _, v := range prompts {
		fmt.Printf("%s.%s\n", v[0], v[1])
	}
	fmt.Println(strings.Repeat("*", 20))
}

func input(pro string) string {
	fmt.Print(pro)
	var txt string
	fmt.Scan(&txt)
	return txt
}

func main() {
	count := 0
	flag := 0
	var password []byte
END:
	for {
		if flag == 0 {
			fmt.Printf("请输入密码:")
			password, _ = gopass.GetPasswd()
		}
		if md5.Sum(password) == md5.Sum([]byte(pass)) {
			flag = 1
			prompt()
			cmd := input("请输入你的操作:")
			if command, ok := commands.Command[cmd]; ok {
				command()
			} else if cmd == "1" {
				break END
			} else {
				fmt.Println("输入指令错误")
			}
		} else {
			count++
			fmt.Printf("密码输入错误！还剩次%d机会\n", 3-count)
			if count == 3 {
				break END
			}
		}
	}
	os.Exit(0)
}
