package main

import (
	"fmt"

	"github.com/creachadair/getpass"

	"user/command"
	"user/cred"
	_ "user/init"
	"user/prompt"
	"user/user"
)

const secret = "42e776142cf72a7a40e56f82fc5e6ce3" //md5 Passwd@1

func main() {
	u :=new(user.Users)
PS:
	passwd, _ := getpass.Prompt("请输入密码:")

	if !cred.Verify(passwd, secret) {
		fmt.Println("密码输入错误,请重新输入密码！")
		goto PS
	} else {
	END:
		for {
			prompt.Prompt()
			cmd := prompt.Input("请输入指令：")
			if cmd == "1" {
				cmd = "add"
			} else if cmd == "2" {
				cmd = "modify"
			} else if cmd == "3" {
				cmd = "delete"
			} else if cmd == "4" {
				cmd = "search"
			}
			if command, ok := command.Commands[cmd]; ok {
				command(u)
			} else if cmd == "5" {
				break END
			} else {
				fmt.Println("指令输入错误！")
			}
		}

	}

}
