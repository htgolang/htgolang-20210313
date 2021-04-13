package handles

import (
	"cmdb/tools"
	"cmdb/user"
	"fmt"
)

var Commands map[string]string = map[string]string{
	"1":    "添加用户",
	"2":    "删除用户",
	"3":    "修改用户",
	"4":    "查询用户信息",
	"5":    "打印用户清单",
	"help": "帮助信息",
}

func handlesRouter(cmd string) {
	switch cmd {
	case "1":
		userData := user.UserInput()
		user.UserAdd(userData)
	case "2":
		userData := tools.StrInput("请输入你要删除的用户ID：")
		user.UserDel(userData)
	case "3":
		userData := tools.StrInput("请输入你要修改的用户ID：")
		user.UserMod(userData)
	case "4":
		userData := tools.StrInput("请输入你要全文检索的信息：")
		user.UserQuery(userData)
	case "5":
		user.UserListPrint()
	case "help":
		tools.HelpPrompt()
	default:
		fmt.Println("请检查您输入的指令！")
	}
}
