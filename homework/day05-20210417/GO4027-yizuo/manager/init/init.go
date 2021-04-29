package init

import (
	"manager/commands"
	"manager/idc"
	"manager/user"
)

func init() {
	commands.Register("添加", user.Add)
	commands.Register("修改", user.Modify)
	commands.Register("删除", user.Delete)
	commands.Register("查询", user.Query)

	commands.Register("添加IDC", idc.Add)
	commands.Register("修改IDC", idc.Modify)
	commands.Register("删除IDC", idc.Delete)
	commands.Register("查询IDC", idc.Query)
}
