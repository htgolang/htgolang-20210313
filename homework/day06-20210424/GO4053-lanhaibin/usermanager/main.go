package main

import (
	"fmt"
	"usermanager/command"
	_ "usermanager/init"
	"usermanager/operation"
	"usermanager/utils"
)

func main() {
	up := operation.UserOperation{}
	up.Login()
	if operation.CurrentUser.Id == 0 {
		return
	}
	for {
		for i, cmd := range command.Commands {
			fmt.Printf("%d.%s\n", i+1, cmd.Desc)
		}
		c := utils.InputInt("请选择操作:")
		if c <= len(command.Commands) {
			cmd := command.Commands[c-1]
			// fmt.Printf("%T", cmd.Cmd())
			cmd.Cmd()
		}
	}
}
