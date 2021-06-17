package init

import (
	"user/commands"
	"user/user"
)

func init() {
	commands.Register("增加", user.Add)
	commands.Register("删除", user.Delete)
	commands.Register("修改", user.Modify)
	commands.Register("查询", user.Query)
}
