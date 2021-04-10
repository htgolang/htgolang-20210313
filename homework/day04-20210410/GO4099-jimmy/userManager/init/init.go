package init

import (
	"userManager/commands"
	"userManager/user"
)

func init() {
	commands.Register("添加", user.Add)
	commands.Register("查询", user.Query)
	commands.Register("修改", user.Edit)
	commands.Register("删除", user.Delete)
}
