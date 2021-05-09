package cmd

import (
	"cmdb/tools"
	"fmt"
)

func Run() {
	// 设置token，验证用户输入的token，失败超过3次退出
	var PASSWORD = "ba3253876aed6bc22d4a6ff53d8406"
	tools.StringInputCheck("请输入进入系统的token: ", PASSWORD, 3)

	// 提示符
	tools.SystemPrompt()

	// 主逻辑
START:
	for {
		cmd := tools.StrInput("请输入要进入的系统:")
		if system := GetSystem(cmd); system != nil {
			system()
		} else if cmd == "quit" {
			break START
		} else {
			fmt.Println("指令错误")
		}
	}

}
