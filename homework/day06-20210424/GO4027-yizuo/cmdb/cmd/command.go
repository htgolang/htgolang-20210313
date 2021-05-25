package cmd

import (
	"cmdb/pkg/admin"
	"cmdb/pkg/todolist"
)

type Callback func()

type Command struct {
	Desc string `json: desc`
	Cmd  Callback
}

var Commands []*Command

func Register(desc string, cmd func()) {
	c := Command{
		Desc: desc,
		Cmd:  cmd,
	}
	Commands = append(Commands, &c)
}

func init() {
	Register("user", admin.AdminSystemRun)
	Register("todolist", todolist.ToDoListSystemRun)
}

func GetSystem(str string) Callback {
	for _, v := range Commands {
		if str == v.Desc {
			return v.Cmd
		}
	}
	return nil
}
