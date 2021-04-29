package handles

import (
	"cmdb/pkg/admin"
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

var user = admin.NewUser()

func HandlesRouter(cmd string) {
	switch cmd {
	case "1":
		user.UserAdd()
	case "2":
		user.UserDel()
	case "3":
		user.UserMod()
	case "4":
		user.UserQuery()
	case "5":
		user.UserListPrint()
	case "help":
		admin.HelpPrompt()
	default:
		fmt.Println("请检查您输入的指令！")
	}
}
