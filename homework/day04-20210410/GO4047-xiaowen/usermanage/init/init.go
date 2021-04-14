package init

import (
	// "usermanage/command"
	"usermanage/tips"
	// "usermanage/user"
)

func init() {
	// command.Comm(user.UserAdd)
	// command.Comm(user.Query)
	// command.Comm(user.DelUser)
	// command.Comm(user.ModifUser)

	tips.Promat("添加用户")
	tips.Promat("用户查询")
	tips.Promat("用户删除")
	tips.Promat("用户修改")
}
