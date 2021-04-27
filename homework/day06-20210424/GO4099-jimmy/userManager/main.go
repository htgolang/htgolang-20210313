package main

import (
	"fmt"
	"strconv"
	"userManager/commands"
	_ "userManager/init"
	"userManager/services"
	"userManager/utils"
)

func main() {
	users := services.Decode()
	if !utils.Auth(users) {
		return
	}
	fmt.Println("欢迎来到用户管理系统")
	for {
		commands.Commands.Prompt()

		cmd, _ := strconv.Atoi(utils.Input("请输入指令: "))
		if callback := commands.Commands.Get(cmd); callback != nil {
			callback()
		} else if cmd == 1 {
			break
		} else {
			fmt.Printf("您输入的指令有误\n")
		}
	}
}
