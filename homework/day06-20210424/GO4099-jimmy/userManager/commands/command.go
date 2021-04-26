package commands

import (
	"fmt"
	"strings"
)

type callback func()

type command struct {
	id       int
	desc     string
	callback callback
}

type commands struct {
	cmds []*command
	next int
}

var Commands = commands{make([]*command, 0), 2}

func (c *commands) Register(desc string, cb callback) {
	c.cmds = append(c.cmds, &command{c.next,desc,cb,})
	c.next++
}

func (c *commands) Prompt() {
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("1. 退出")
	for _, cmd := range c.cmds {
		fmt.Printf("%d. %s\n", cmd.id, cmd.desc)
	}
	fmt.Println(strings.Repeat("=", 20))
}

func (c *commands) Get(id int) callback {
	for _, c := range c.cmds {
		if c.id == id {
			return c.callback
		}
	}
	return nil
}

func Register(desc string, cb callback) {
	Commands.Register(desc, cb)
}

