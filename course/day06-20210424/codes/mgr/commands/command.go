package commands

import "fmt"

var id int = 2

type Callback func()

type Command struct {
	ID       int
	Desc     string
	Callback Callback
}

var Commands = make(commands, 0)

type commands []*Command

func (commands) Get(id int) Callback {
	for _, cmd := range Commands {
		if id == cmd.ID {
			return cmd.Callback
		}
	}
	return nil
}

func (commands) Prompts(tpl string) {
	for _, cmd := range Commands {
		fmt.Printf(tpl, cmd.ID, cmd.Desc)
	}
}

func Register(desc string, callback Callback) {
	Commands = append(Commands, &Command{id, desc, callback})
	id++
}
