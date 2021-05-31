package command

import (
	"fmt"
	"strings"
)

type command struct {
	id       int
	desc     string
	callback func()
}

type Commands struct {
	cmds []command
	next int
}

var UserCommands = Commands{[]command{}, 2}
var TaskCommands = Commands{[]command{}, 2}
var TotalCommands = map[int]*Commands{2: &UserCommands, 3: &TaskCommands}

func (c *Commands) Register(desc string, callback func()) {
	c.cmds = append(c.cmds, command{id: c.next, desc: desc, callback: callback})
	c.next++
}

// func (c Commands) Prompt()  {
// 	fmt.Println("1 返回上一级")
// 	for _, v := range c.cmds {
// 		fmt.Printf("%d %s\n", v.id, v.desc)
// 	}
// }

func (c Commands) Prompt()  {
	fmt.Println(strings.Repeat("*", 15))
	fmt.Println("1 返回上一级")
	for _, v := range c.cmds {
		fmt.Printf("%d %s\n", v.id, v.desc)
	}
	fmt.Println(strings.Repeat("*", 15))
}

func (c Commands) Get(id int) func() {
	for _, v := range c.cmds {
		if id == v.id {
			return v.callback
		}
	}
	return nil
}

func Register(c *Commands, desc string, callback func()) {
	c.Register(desc, callback)
}
