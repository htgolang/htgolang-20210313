package admin

import "cmdb/pkg/todolist"

type callback func()

type Command struct {
	Desc string `json: desc`
	Cmd  callback
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
	Register("csv", NewCsvClient)
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
