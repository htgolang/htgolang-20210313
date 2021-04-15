package init

import (
	"users/commands"
	"users/idc"
	"users/user"
)

func init() {
	commands.Register("add",user.Add)
}