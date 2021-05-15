package commands

import (
	"fmt"
	"mgr/model/user"
)

//Define callback
type callback func([]*user.User)

//define command structure
type command struct {
	id int
	desc string
	callback callback
}

//define commands structure
type commands struct {
	cmds []command
	next int
}

//init command
var Commands = commands{make([]command, 0, 10), 2}

//register command
func (cs *commands) Register(desc string, callback callback){
	cs.cmds = append(cs.cmds, command{cs.next, desc, callback})
	cs.next++
}

//menu
func (cs commands) Prompt(start int, tpl string)  {
	for _,cmd:= range cs.cmds{
		fmt.Printf(tpl, cmd.id, cmd.desc)
	}
}

//get callback
func (cs commands) Get(id int) callback  {
	for _, cmd := range cs.cmds{
		if cmd.id == id{
			return cmd.callback
		}
	}
	return nil
}

//get all callback
func (cs commands) GetAll() {
	for _, cmd := range cs.cmds {
		fmt.Println(cmd.id,cmd.desc,cmd.callback)
	}
}

//register
func Register(desc string, callback callback) {
	Commands.Register(desc, callback)
}
