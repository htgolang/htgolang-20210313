package init

import (
	"userManager/commands"
	"userManager/services"
)

func init() {
	var userService services.UserService
	var persistenceService services.PersistenceService
	commands.Register("添加", userService.Add)
	commands.Register("删除", userService.Delete)
	commands.Register("修改", userService.Modify)
	commands.Register("查询", userService.Query)
	commands.Register("修改密码", userService.ChangePassword)
	commands.Register("持久化", persistenceService.Pst)
}
