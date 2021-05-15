package init

import (
	// "usermanage/command"
	"usermanage/tips"
	// "usermanage/user"
	// "usermanage/todo"
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



	tips.TaskPro("任务增加")
	tips.TaskPro("任务删除")
	tips.TaskPro("任务修改")
	tips.TaskPro("任务查询")
}
