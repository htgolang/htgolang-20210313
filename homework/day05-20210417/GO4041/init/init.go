package init

import (
	"user/command"
	"user/user"
)

func init() {
	command.Register("add", user.Add)
	command.Register("modify", user.Modify)
	command.Register("delete", user.Delete)
	command.Register("search", user.Search)
}
