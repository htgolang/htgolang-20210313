package init

import (
	"cmdb/tools"
	"cmdb/user"
)

func init() {
	// 初始化用户
	initUserList()
}

func initUserList() {
	userMapData := map[string]string{
		"Id":     "1",
		"Name":   "yizuo",
		"Passwd": tools.Md5sum("yizuo"),
		"Addr":   "Wuhan",
		"Tel":    "66666666",
	}
	user.UserList = append(user.UserList, userMapData)
}
