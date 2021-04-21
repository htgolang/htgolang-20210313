package usermanage

import (
	"fmt"
	"strconv"
	"user/utils"
	"strings"
)

var UserInfo = []map[string]string{
	{"id": "1", "name": "wx", "addr": "北京", "tel": "135467"},
	{"id": "3", "name": "jack", "addr": "四川", "tel": "152467"},
	{"id": "4", "name": "jim", "addr": "河北", "tel": "135167"},
}

//生成用户ID
func generateID() string {
	id := 0
	for _, v := range UserInfo {
		i, _ := strconv.Atoi(v["id"])
		if i > id {
			id = i
		}
	}
	return strconv.Itoa(id + 1)
}

func Add() {
	name := utils.Input("请输入名字:")
	addr := utils.Input("请输入地址:")
	tel := utils.Input("请输入电话:")

	tmpuser := map[string]string{"id": generateID(), "name": name, "addr": addr, "tel": tel}
	UserInfo = append(UserInfo, tmpuser)

	fmt.Println(UserInfo)

}

func Delete() {
	id := utils.Input("请输入要删除的用户id:")
	index := utils.CheckUserExist(id, UserInfo)
	if index != -1 {
		fmt.Println(UserInfo[index])
		confirm := utils.Input("是否确认删除(y/n)?")
		switch confirm {
		case "y", "Y":
			UserInfo = append(UserInfo[:index], UserInfo[index+1:]...)
			fmt.Println(UserInfo)
		case "n":
			fmt.Println("取消删除")
		default:
			fmt.Println("输入错误")
		}
	} else {
		fmt.Println("用户不存在")
	}
}

func Modify() {
	id := utils.Input("请输入要修改的用户id:")
	index := utils.CheckUserExist(id, UserInfo)

	if index != -1{
		fmt.Println(UserInfo[index])
		confirm := utils.Input("是否确认修改(y/n)?")
		switch confirm {
		case "y", "Y":
			newName := utils.Input("请输入新的用户名:")
			newAddr := utils.Input("请输入新的地址:")
			newTel := utils.Input("请输入新的电话:")
			UserInfo[index]["name"] = newName
			UserInfo[index]["addr"] = newAddr
			UserInfo[index]["tel"] = newTel
			fmt.Println(UserInfo)		
		case "n":
			fmt.Println("取消修改")
		default:
			fmt.Println("输入错误")
		}
	} else {
		fmt.Println("用户不存在")
	}

}

func Query() {
	keyWord := utils.Input("请输入要查询的关键字:")

	for _,v := range UserInfo{
		//id不进行关键字匹配
		if strings.Contains(v["name"], keyWord) || strings.Contains(v["addr"], keyWord) || strings.Contains(v["tel"], keyWord) {
			fmt.Println(v)
		}
	}
}
