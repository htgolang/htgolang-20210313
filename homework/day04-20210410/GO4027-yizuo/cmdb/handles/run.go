package handles

import (
	_ "cmdb/init"
	"cmdb/tools"
	"fmt"
)

func Run() {
	// 用户密码登录检测，超过3次退出
	if !HandlerUserLoginAuth() {
		fmt.Println("密码错误超过3次，已退出！")
		return
	}

	// 提示符
	tools.HelpPrompt()

	// 主逻辑
START:
	for {
		cmd := tools.StrInput("请输入指令:")
		if _, ok := Commands[cmd]; ok {
			handlesRouter(cmd)
		} else if cmd == "0" {
			break START
		} else {
			fmt.Println("指令错误")
		}
	}

}
