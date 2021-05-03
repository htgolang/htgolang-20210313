package init

import (
	"os"
	"usermanager/command"
	"usermanager/operation"
)

func init() {
	up := operation.UserOperation{}
	command.Register("添加用户", up.Add)
	command.Register("删除用户", up.Delete)
	command.Register("查询用户", up.Query)
	command.Register("修改用户", up.Modify)

	tp := operation.TodoOperation{}
	command.Register("添加任务", tp.Add)
	command.Register("删除任务", tp.Delete)
	command.Register("查询任务", tp.Query)
	command.Register("修改任务", tp.Modify)

	command.Register("退出", func() {
		os.Exit(0)
	})
}
