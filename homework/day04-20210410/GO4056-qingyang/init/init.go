package init

import (
	"fmt"
	"user/common"
	"user/utils"
)

func init() {
	// 初始化用户
	initUserList()
	fmt.Println("执行了init的init")
	fmt.Println(common.UserList)
}

func initUserList() {
	userMapData := map[string]string{
		"Id":       string(common.Id),
		"Name":     "qingyang",
		"PassWord": utils.Md5sum("qingyang"),
		"Addr":     "shenzhen",
		"Tel":      "11111111",
	}
	common.UserList = append(common.UserList, userMapData)
}
