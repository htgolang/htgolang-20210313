package usermanage

import (
	"fmt"
	"strings"
	"user/module"
	"user/utils"

	"github.com/princebot/getpass"
)

const (
	user     = "admin"
	password = "a3dcb8aae0d84122a1778e26da285857" //123456
)

//生成用户ID
func generateID() int {
	id := 0
	for _, v := range module.UserInfo {
		i := v.Id
		if i > id {
			id = i
		}
	}
	return id + 1
}

//用户登录
func Login() bool {
	for i := 0; i < 3; i++ {
		name := utils.Input("请输入用户名:")
		pw, _ := getpass.Get("请输入密码:")

		if name == "admin" && utils.Md5Encrypt(pw) == password {
			fmt.Printf("登录成功!\n\n")
			return true
		} else {
			fmt.Printf("用户名或密码错误, 还可尝试%d次\n\n", 2-i)
		}
	}
	return false
}

func Add() {
	name := utils.Input("请输入名字:")
	addr := utils.Input("请输入地址:")
	tel := utils.Input("请输入电话:")

	tmpuser := module.User{Id: generateID(), Name: name, Addr: addr, Tel: tel}
	module.UserInfo = append(module.UserInfo, tmpuser)

	fmt.Println(module.UserInfo)

}

func Delete() {
	id := utils.Input("请输入要删除的用户id:")
	index := utils.CheckUserExist(id, module.UserInfo)
	if index != -1 {
		fmt.Println(module.UserInfo[index])
		confirm := utils.Input("是否确认删除(y/n)?")
		switch confirm {
		case "y", "Y":
			module.UserInfo = append(module.UserInfo[:index], module.UserInfo[index+1:]...)
			fmt.Println(module.UserInfo)
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
	index := utils.CheckUserExist(id, module.UserInfo)

	if index != -1 {
		fmt.Println(module.UserInfo[index])
		confirm := utils.Input("是否确认修改(y/n)?")
		switch confirm {
		case "y", "Y":
			newName := utils.Input("请输入新的用户名:")
			newAddr := utils.Input("请输入新的地址:")
			newTel := utils.Input("请输入新的电话:")
			module.UserInfo[index].Name = newName
			module.UserInfo[index].Addr = newAddr
			module.UserInfo[index].Tel = newTel
			fmt.Println(module.UserInfo)
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

	flag := false
	for _, v := range module.UserInfo {
		//id不进行关键字匹配
		if strings.Contains(v.Name, keyWord) || strings.Contains(v.Addr, keyWord) || strings.Contains(v.Tel, keyWord) {
			fmt.Println(v)
			flag = true
		}
	}

	if !flag {
		fmt.Println("没有匹配用户")
	}

}
