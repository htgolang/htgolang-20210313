package init

import (
	"os"
	"usermanager/command"
	"usermanager/operation"
	"usermanager/service"
)

func init() {
	up := operation.UserOperation{}
	command.Register("添加用户", up.Add)
	command.Register("删除用户", up.Delete)
	command.Register("查询用户", up.Query)
	command.Register("修改用户", up.Modify)
	command.Register("修改密码", up.ChangePassword)

	tp := operation.TodoOperation{}
	command.Register("添加任务", tp.Add)
	command.Register("删除任务", tp.Delete)
	command.Register("查询任务", tp.Query)
	command.Register("修改任务", tp.Modify)

	ds := operation.DbOperation{}
	command.Register("修改序列化", ds.ChangeDb)

	command.Register("退出", func() {
		service.Db.EncodeUser()
		os.Exit(0)
	})

	service.Db = service.DbService{Type: "gob"}
	service.Db.DecodeUser()
}
