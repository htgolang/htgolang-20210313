package common

import (
	"crypto/md5"
	"fmt"
	"github.com/princebot/getpass"
	"strings"
	"userManager/commands"
)

const (
	MaxAuthCount = 3
	Password = "e10adc3949ba59abbe56e057f20f883e"
)

// 密码验证
func Auth() bool {
	for i:=0; i< MaxAuthCount; i++ {
		pwd, err := getpass.Get("请输入密码: ")
		if err != nil {
			panic(err)
		}
		if fmt.Sprintf("%x", md5.Sum(pwd)) == Password {
			return true
		} else {
			if i == 2 {
				fmt.Println("密码错误，您的三次机会已用完")
			} else {
				fmt.Printf("密码错误，您还有%d次机会\n", MaxAuthCount - i - 1)
			}
		}
	}
	return false
}

// 打印提示
func Prompts() {
	p := commands.Prompts
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("1 退出")
	for _, v := range p {
		fmt.Println(v[0], v[1])
	}
	fmt.Println(strings.Repeat("=", 20))
}

// 输入内容
func Input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}

func InputUser() map[string]string {
	name := Input("请输入用户名: ")
	tel := Input("请输入电话: ")
	addr := Input("请输入地址: ")
	user := map[string]string{}
	user["name"] = name
	user["tel"] = tel
	user["addr"] = addr
	return user
}
