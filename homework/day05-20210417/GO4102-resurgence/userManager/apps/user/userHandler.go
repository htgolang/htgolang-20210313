package user

import (
	"fmt"
	. "userManager/config"
	"userManager/utils"
)

func GetUserList(userMap map[string]*User) {
	if lens := len(userMap); lens == 0 {
		fmt.Println("目前还没有用户")
		return
	}

	for id, user := range userMap {
		fmt.Printf("用户id: %s , 姓名: %s , 手机号: %s\n", id, user.Name, user.Phone)
	}
}

func AddUser(userMap map[string]*User) {
	name := utils.Input("请输入用户名: ")
	phone := utils.Input("请输入手机号: ")

	fmt.Println(name, phone)
	user, id := NewUser(name, phone)
	userMap[id] = user
	fmt.Printf("%v\n", user)
}

func UpdateUser(userMap map[string]*User) {
	upID := utils.Input("输入要更新的用户id: ")
	user, ok := userMap[upID]
	if !ok {
		fmt.Println("查无此人")
		return
	}

	attr := utils.Input("输入要修改的属性: ")
	val := utils.Input("输入要修改的值: ")

	switch attr {
	case "name":
		user.Name = val
	case "phone":
		user.Phone = val
	default:
		fmt.Println("输入非法, 请输入 name 或 phone 中的属性")
		return
	}
	fmt.Printf("用户 %s 更新成功\n", upID)
}

func DeleteUser(userMap map[string]*User) {
	delID := utils.Input("输入要删除的用户id: ")

	_, ok := userMap[delID]
	if !ok {
		fmt.Println("查无此人")
		return
	}

	delete(userMap, delID)
	if _, ok := userMap[delID]; !ok {
		fmt.Printf("用户 %s 删除成功", delID)
	}
}
