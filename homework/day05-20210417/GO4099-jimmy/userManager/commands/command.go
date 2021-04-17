package commands

import "fmt"

type callback func()

type command struct {
	id   int
	desc string
	cb   callback
}

type commands struct {
	cmds []command
	next int
}

var Commands = commands{make([]command, 0, 10), 2}

func (c *commands) Register(desc string, cb callback) {
	c.cmds = append(c.cmds, command{id: c.next, desc: desc, cb: cb})
	c.next++
}

func (c *commands) Prompts() {
	for _, cmd := range c.cmds {
		fmt.Printf("%d. %s\n", cmd.id, cmd.desc)
	}
}

func (c *commands) Get(id int) callback {
	for _, cmd := range c.cmds {
		if cmd.id == id {
			return cmd.cb
		}
	}
	return nil
}

func Register(desc string, cb callback) {
	Commands.Register(desc, cb)
}
