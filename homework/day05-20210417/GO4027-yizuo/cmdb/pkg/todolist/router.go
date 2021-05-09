package todolist

import (
	"cmdb/tools"
	"fmt"
)

type Callback func()

type todoListCommand struct {
	Desc string `json: desc`
	Cmd  Callback
}

var todoListCommands []todoListCommand

func Register(desc string, cmd Callback) {
	c := todoListCommand{
		Desc: desc,
		Cmd:  cmd,
	}
	todoListCommands = append(todoListCommands, c)
}

func getToDoListSystem(str string) Callback {
	for _, v := range todoListCommands {
		if str == v.Desc {
			return v.Cmd
		}
	}
	return nil
}

func toDoListRoutersInit() {
	// 注册用户系统路由
	Register("add", TodoListAdd)
	Register("del", TodoListDel)
	Register("mod", TodoListMod)
	Register("query", TodoListQuery)
	Register("list", TodoListListPrint)
	Register("help", tools.ToDoListHelpPrompt)
}

func ToDoListSystemRun() {
	// 初始化用户系统路由
	toDoListRoutersInit()

	// 提示符
	tools.ToDoListHelpPrompt()

	// 主逻辑
START:
	for {
		cmd := tools.StrInput("请输入指令:")
		if system := getToDoListSystem(cmd); system != nil {
			system()
		} else if cmd == "quit" {
			tools.SystemPrompt()
			break START
		} else {
			fmt.Println("指令错误")
		}
	}

}
