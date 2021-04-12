package main

import (
	"fmt"

	"github.com/creachadair/getpass"

	"user/command"
	"user/cred"
	_ "user/init"
	"user/prompt"
)

const secret = "42e776142cf72a7a40e56f82fc5e6ce3" //md5 Passwd@1

func main() {
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
			if command, ok := command.Commands[cmd]; ok {
				command()
			} else if cmd == "5" {
				break END
			} else {
				fmt.Println("指令输入错误！")
			}
		}

	}

}
