package user

import (
	"fmt"
	"strconv"
	"user/common"
)

func AddUser() {
	var (
		name, addr, tel string
	)
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入地址：")
	fmt.Scanln(&addr)
	fmt.Println("请输入电话：")
	fmt.Scanln(&tel)
	common.Id++
	common.UserList = append(common.UserList, map[string]string{
		"id":   strconv.Itoa(common.Id),
		"name": name,
		"addr": addr,
		"tel":  tel,
	})
	fmt.Println("添加用户成功")
}
