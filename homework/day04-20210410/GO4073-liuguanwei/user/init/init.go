package init

import (
	"user/user"
	"user/commands"
)

func init(){
	commands.Register("add", user.Add)
	commands.Register("edit", user.Edit)
	commands.Register("del", user.Del)
	commands.Register("query", user.Query)

}