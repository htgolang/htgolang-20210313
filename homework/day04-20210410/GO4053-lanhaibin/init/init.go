package init

import (
	"os"
	"user/command"
	"user/user"
)

func init() {
	command.Register(user.AddAction, "添加")
	command.Register(user.DeleteAction, "删除")
	command.Register(user.QueryAction, "查询")
	command.Register(user.ModifyAction, "修改")
	command.Register(func() { os.Exit(0) }, "退出")
}
