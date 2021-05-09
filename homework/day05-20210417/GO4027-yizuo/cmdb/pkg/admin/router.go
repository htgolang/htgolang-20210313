package admin

import (
	"cmdb/tools"
	"fmt"
)

type Callback func()

type userCommand struct {
	Desc string `json: desc`
	Cmd  Callback
}

var userCommands []userCommand

func Register(desc string, cmd Callback) {
	c := userCommand{
		Desc: desc,
		Cmd:  cmd,
	}
	userCommands = append(userCommands, c)
}

func getAdminSystem(str string) Callback {
	for _, v := range userCommands {
		if str == v.Desc {
			return v.Cmd
		}
	}
	return nil
}

func adminRoutersInit() {
	// 初始化第一个用户
	var user = NewUser("1", "yizuo", tools.Md5sum("yizuo"), "Wuhan", "66666666")
	UserList = append(UserList, user)

	// 注册用户系统路由
	Register("add", user.UserAdd)
	Register("del", user.UserDel)
	Register("mod", user.UserMod)
	Register("query", user.UserQuery)
	Register("list", user.UserListPrint)
	Register("help", tools.AdminHelpPrompt)
}

func AdminSystemRun() {
	// 初始化用户系统路由
	adminRoutersInit()

	// 用户密码登录检测，超过3次退出
	if !HandlerUserLoginAuth() {
		fmt.Println("密码错误超过3次，已退出！")
		return
	}

	// 提示符
	tools.AdminHelpPrompt()

	// 主逻辑
START:
	for {
		cmd := tools.StrInput("请输入指令:")
		if system := getAdminSystem(cmd); system != nil {
			system()
		} else if cmd == "quit" {
			tools.SystemPrompt()
			break START
		} else {
			fmt.Println("指令错误")
		}
	}

}
