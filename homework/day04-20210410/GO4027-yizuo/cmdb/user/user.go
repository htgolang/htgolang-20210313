package user

import (
	"cmdb/tools"
	"fmt"
	"strconv"
	"strings"
)

var UserList []map[string]string

// 增加用户
func UserAdd(userAddData map[string]string) {
	UserList = append(UserList, userAddData)
}

// 删除用户
func UserDel(userDelId string) {
	userDelPlace, ok := FindIdExists(userDelId)
	if !ok {
		fmt.Println("你输入的用户不存在！")
		return
	}

	fmt.Printf("你要删除的的用户数据为：%v\n", UserList[userDelPlace])
	if !tools.UserInputExecuted("确认要删除用户吗？：(yes/no ; default no)") {
		fmt.Println("输入错误，返回首页。")
		return
	}

	UserList = append(UserList[:userDelPlace], UserList[userDelPlace+1:]...)
	fmt.Println(UserList)

}

// 修改用户
func UserMod(userModId string) {
	userModPlace, ok := FindIdExists(userModId)
	if !ok {
		fmt.Println("你输入的用户不存在！")
		return
	}

	fmt.Printf("你要修改的的用户数据为：%v\n", UserList[userModPlace])
	if !tools.UserInputExecuted("确认要修改用户数据吗？：(yes/no ; default no)") {
		fmt.Println("输入错误，返回首页。")
		return
	}

	UserList[userModPlace] = UserInput()
	fmt.Println(UserList)
}

// 检索数据返回
func UserQuery(userQueryData string) {
	userQueryPlace, ok := FindStrDataExists(userQueryData)
	if !ok {
		fmt.Println("未查询到你输入的信息。")
		return
	}
	fmt.Printf("查询到的用户数据为：\n%v\n", UserList[userQueryPlace])
}

// 打印当前所有数据
func UserListPrint() {
	fmt.Println(UserList)
}

// 查找用户ID是否存在，如果存在返回对应切片位置并返回true，不存在返回false。
func FindIdExists(data string) (int, bool) {
	for k, v := range UserList {
		if v["Id"] == data {
			return k, true
		}
	}
	return 0, false
}

// 全文检索传入字符串是否存在，如果存在返回对应切片位置并返回true，不存在返回false。
func FindStrDataExists(data string) (int, bool) {
	for place, mapData := range UserList {
		for _, v := range mapData {
			if strings.Contains(v, data) {
				return place, true
			}
		}
	}
	return 0, false
}

// 查找用户ID最大的值+1
func FindMaxIdNum() (maxIdNum int) {
	maxIdNum = 0
	for _, v := range UserList {
		idNum, _ := strconv.Atoi(v["Id"])
		if idNum > maxIdNum {
			maxIdNum = idNum + 1

		}
	}
	return maxIdNum
}

// 获取用户输入，返回用户输入的map合集
func UserInput() (userInputData map[string]string) {
	var id, name, passwd, addr, tel string
	id = strconv.Itoa(FindMaxIdNum())
	name = tools.StrInput("请输入用户名称：")
	passwd = tools.StrInput("请输入用户密码：")
	addr = tools.StrInput("请输入用户地址：")
	tel = tools.StrInput("请输入联系方式：")

	userInputData = map[string]string{
		"Id":     id,
		"Name":   name,
		"Passwd": tools.Md5sum(passwd),
		"Addr":   addr,
		"Tel":    tel,
	}
	return userInputData
}

// 用户登录信息检测
func UserLoginAuth(userData, passwdData string) bool {
	for _, v := range UserList {
		if v["Name"] == userData && v["Passwd"] == tools.Md5sum(passwdData) {
			return true
		}
	}
	return false
}
