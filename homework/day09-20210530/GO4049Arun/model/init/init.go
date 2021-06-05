package init

import (
	"mgr/controller/commands"
	//"mgr/model/user"
	"mgr/controller/user"
)

func init() {
	commands.Register("新增用户", user.Add)
	commands.Register("修改用户", user.ModifyUser)
	commands.Register("删除用户)", user.DeleteUser)
	commands.Register("查询用户信息", user.FindUser)

	//commands.Register("添加IDC", idc.Add)
	//commands.Register("修改IDC", idc.Modify)
	//commands.Register("删除IDC", idc.Delete)
	//commands.Register("查询IDC", idc.Query)
}
