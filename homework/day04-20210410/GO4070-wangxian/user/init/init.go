package init

import (
	"user/command"
	"user/usermanage"
)

func init() {
	command.Register("添加", usermanage.Add)
	command.Register("删除", usermanage.Delete)
	command.Register("修改", usermanage.Modify)
	command.Register("查询", usermanage.Query)
}
