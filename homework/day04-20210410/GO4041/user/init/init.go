package init

import (
	"user/command"
	"user/user"
)

func init() {
	command.Register("添加", user.Add)
	command.Register("修改", user.Modify)
	command.Register("删除", user.Delete)
	command.Register("查询", user.Search)
}
