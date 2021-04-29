package commands

import (
	"fmt"
)

// 定义callback类型
type callback func()

// 定义command结构体
type command struct {
	id       int
	desc     string
	callback callback
}

// 定义commands结构体
type commands struct {
	cmds []command
	next int
}

// 初始化Commands
var Commands = commands{make([]command, 0, 10), 2}

// 注册command
func (cs *commands) Register(desc string, callback callback) {
	cs.cmds = append(cs.cmds, command{cs.next, desc, callback})
	cs.next++
}

// 菜单
func (cs commands) Prompt(start int, tpl string) {
	for _, cmd := range cs.cmds {
		fmt.Printf(tpl, cmd.id, cmd.desc)
	}
}

// 获取callback
func (cs commands) Get(id int) callback {
	for _, cmd := range cs.cmds {
		if cmd.id == id {
			return cmd.callback
		}
	}
	return nil
}

// 注册
func Register(desc string, callback callback) {
	Commands.Register(desc, callback)
}
