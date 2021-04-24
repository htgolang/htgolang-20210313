package init

import (
	"mgr/commands"
	"mgr/services"
)

func init() {
	var userService services.UserService
	var idcService services.IDCService

	commands.Register("添加", userService.Add)
	commands.Register("修改", userService.Modify)
	commands.Register("删除", userService.Delete)
	commands.Register("查询", userService.Query)

	commands.Register("添加IDC", idcService.Add)
	commands.Register("修改IDC", idcService.Modify)
	commands.Register("删除IDC", idcService.Delete)
	commands.Register("查询IDC", idcService.Query)
}
