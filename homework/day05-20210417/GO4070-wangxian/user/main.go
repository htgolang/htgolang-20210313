package main

import (
	"fmt"
	"strconv"
	"strings"
	"user/command"
	_ "user/init"
	"user/usermanage"
	"user/utils"
)

// var commands = command.Commands

//打印提示信息
func prompt(c *command.Commands) {
	fmt.Println(strings.Repeat("*", 15))
	c.Prompt()
	fmt.Println(strings.Repeat("*", 15))
}

//菜单
func menu(c *command.Commands) {
	for {
		prompt(c)
		cmd, _ := strconv.Atoi(utils.Input("请输入指令:"))
		if f := c.Get(cmd); f != nil {
			f()
		} else if cmd == 1 {
			return
		} else {
			fmt.Println("指令错误")
		}
	}
}

func main() {
	if usermanage.Login() {
		fmt.Println("欢迎使用xxx管理系统")
	END:
		for {
			fmt.Println(`
1 退出
2 用户管理
3 任务管理`)
			cmd, _ := strconv.Atoi(utils.Input("请输入指令:"))

			if v, ok := command.TotalCommands[cmd]; ok {
				// fmt.Println(v)
				menu(v)
			} else if cmd == 1 {
				break END
			} else {
				fmt.Println("没有此功能")
			}
		}
	} else {
		fmt.Println("登录失败，请稍后再试!")
	}

}
