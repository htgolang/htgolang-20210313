package admin

import (
	"cmdb/tools"
	"fmt"
)

type Callback func()

type userCommand struct {
	Desc string
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
	// var user = NewUser("1", "yizuo", tools.Md5sum("yizuo"), "Wuhan", "66666666")
	// UserList = append(UserList, user)
	// 读取用户列表
	ReadUsersDataFromCsv()

	// 注册用户系统路由
	var u *Users
	Register("add", u.UserAdd)
	Register("del", u.UserDel)
	Register("mod", u.UserMod)
	Register("query", u.UserQuery)
	Register("list", u.UserListPrint)
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
			// 备份用户数据文件
			CopyUserDataFile()
			// 检查持久化的文件数量
			PersistenceOfLastNChanges()
			// 打印之前的用户系统提示
			tools.SystemPrompt()
			break START
		} else {
			fmt.Println("指令错误")
		}
	}

}
