package main

import (
	"fmt"
	"userManager/commands"
	"userManager/common"
	_ "userManager/init"
)

func main() {
	if !common.Auth() {
		return
	}
	fmt.Println("登录成功，欢迎来到用户管理系统")

	for {
		common.Prompts()
		instruction := common.Input("请输入指令: ")
		if callback, ok := commands.Commands[instruction]; ok {
			callback()
		} else if instruction == "1" {
			break
		} else {
			fmt.Printf("您输入的指令有误\n")
		}
	}
}
