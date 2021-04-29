package cmd

import (
	"cmdb/handles"
	"cmdb/pkg/admin"
	"fmt"
)

func Run() {
	// 用户密码登录检测，超过3次退出
	if !admin.HandlerUserLoginAuth() {
		fmt.Println("密码错误超过3次，已退出！")
		return
	}

	// 提示符
	admin.HelpPrompt()

	// 主逻辑
START:
	for {
		cmd := admin.StrInput("请输入指令:")
		if _, ok := handles.Commands[cmd]; ok {
			handles.HandlesRouter(cmd)
		} else if cmd == "0" {
			break START
		} else {
			fmt.Println("指令错误")
		}
	}

}
