package main

import (
	"fmt"
	"usermanager/command"
	"usermanager/common"
	_ "usermanager/init"
	"usermanager/utils"
)

func main() {
	if !common.Login() {
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
