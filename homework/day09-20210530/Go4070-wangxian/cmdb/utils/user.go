package utils

import (
	"cmdb/modules"
)

//检查用户名是否冲突
func CheckUserName(name string, users []modules.User) bool {
	for _, v := range users {
		if name == v.Name {
			return true
		}
	}
	return false
}

//检查用户是否存在,存在返回在切片中的索引，不存在返回-1
func CheckUserExist(uid int, users []modules.User) int {
	for i, v := range users {
		if uid == v.Id {
			return i
		}
	}
	return -1
}

//检查号码是否冲突
func CheckTel(tel string, users []modules.User) bool {
	for _, v := range users {
		if tel == v.Tel {
			return true
		}
	}
	return false
}
