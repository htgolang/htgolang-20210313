package init

import (
	"userManager/commands"
	"userManager/todo"
	"userManager/user"
)

func init() {
	commands.Register("添加用户", user.Add)
	commands.Register("查询用户", user.Query)
	commands.Register("修改用户", user.Modify)
	commands.Register("删除用户", user.Delete)

	commands.Register("添加任务", todo.Add)
	commands.Register("查询任务", todo.Query)
	commands.Register("修改任务", todo.Modify)
	commands.Register("删除任务", todo.Delete)
}
