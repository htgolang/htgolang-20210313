package common

import (
	"usermanager/model"
	"usermanager/utils"
)

func InputUser() model.User {
	user := model.User{}
	user.Name = utils.Input("请输入用户名:")
	user.Addr = utils.Input("请输入地址:")
	user.Tel = utils.Input("请输入电话号码:")
	return user
}
