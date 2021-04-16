package user

import (
	"fmt"
	"strconv"
	"strings"
	"user/common"
)

func DeleteUser() {
	var id int
	fmt.Println("请输入要删除的用户id：")
	fmt.Scanln(&id)
	var flag string
	for i, v := range common.UserList {
		val, _ := strconv.Atoi(v["id"])
		if val == id+1 {
			fmt.Println("用户存在，是否删除(y/yes/Y/YES)")
			fmt.Scanln(&flag)
			if strings.EqualFold(flag, "y") || strings.EqualFold(flag, "yes") || strings.EqualFold(flag, "Y") || strings.EqualFold(flag, "YES") {
				common.UserList = append(common.UserList[:i], common.UserList[i+1:]...)
			} else {
				break
			}
		} else {
			fmt.Println("不存在该用户")
			break
		}
	}
}
