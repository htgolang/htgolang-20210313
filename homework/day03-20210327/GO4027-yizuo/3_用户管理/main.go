package main

import (
	"fmt"
	"strconv"
	"strings"
)

var userList = make([]map[string]string, 0)

// 用户初始化
func init() {
	userMapData := map[string]string{
		"id":   "1",
		"name": "yizuo",
		"addr": "Wuhan",
		"tel":  "66666666",
	}
	userList = append(userList, userMapData)
}

// 增加用户
func userAdd(userAddData map[string]string) {
	userList = append(userList, userAddData)
}

// 删除用户
func userDel(userDelId string) {
	userDelPlace, ok := findIdExists(userDelId)
	if !ok {
		fmt.Println("你输入的用户不存在！")
		return
	}

	fmt.Printf("你要删除的的用户数据为：%v\n", userList[userDelPlace])
	if !userInputExecuted("确认要删除用户吗？：(yes/no ; default no)") {
		fmt.Println("输入错误，返回首页。")
		return
	}

	userList = append(userList[:userDelPlace], userList[userDelPlace+1:]...)
	fmt.Println(userList)

}

// 修改用户
func userMod(userModId string) {
	userModPlace, ok := findIdExists(userModId)
	if !ok {
		fmt.Println("你输入的用户不存在！")
		return
	}

	fmt.Printf("你要修改的的用户数据为：%v\n", userList[userModPlace])
	if !userInputExecuted("确认要修改用户数据吗？：(yes/no ; default no)") {
		fmt.Println("输入错误，返回首页。")
		return
	}

	userList[userModPlace] = userInput()
	fmt.Println(userList)
}

// 检索数据返回
func userQuery(userQueryData string) {
	userQueryPlace, ok := findStrDataExists(userQueryData)
	if !ok {
		fmt.Println("未查询到你输入的信息。")
		return
	}
	fmt.Printf("查询到的用户数据为：\n%v\n", userList[userQueryPlace])
}

// 查找用户ID是否存在，如果存在返回对应切片位置并返回true，不存在返回false。
func findIdExists(data string) (int, bool) {
	for k, v := range userList {
		if v["id"] == data {
			return k, true
		}
	}
	return 0, false
}

// 全文检索传入字符串是否存在，如果存在返回对应切片位置并返回true，不存在返回false。
func findStrDataExists(data string) (int, bool) {
	for place, mapData := range userList {
		for _, v := range mapData {
			if strings.Contains(v, data) {
				return place, true
			}
		}
	}
	return 0, false
}

// 确认是否执行
func userInputExecuted(data string) bool {
	userInputExecutedData := strInput(data)
	switch userInputExecutedData {
	case "y", "Y", "yes", "YES":
		return true
	default:
		return false
	}
}

// 获取用户输入，返回用户输入的map合集
func userInput() (userInputData map[string]string) {
	var id, name, addr, tel string
	id = strconv.Itoa(findMaxIdNum())
	name = strInput("请输入用户名称：")
	addr = strInput("请输入用户地址：")
	tel = strInput("请输入联系方式：")

	userInputData = map[string]string{
		"id":   id,
		"name": name,
		"addr": addr,
		"tel":  tel,
	}
	return userInputData
}

// 传入字符串，自动返回用户输入的值
func strInput(strData string) (strInput string) {
	fmt.Println(strData)
	fmt.Scan(&strInput)
	return strInput
}

// 查找用户ID最大的值+1
func findMaxIdNum() (maxIdNum int) {
	maxIdNum = 0
	for _, v := range userList {
		idNum, _ := strconv.Atoi(v["id"])
		if idNum > maxIdNum {
			maxIdNum = idNum + 1
		}
	}
	return maxIdNum
}

// 主逻辑
func main() {
END:
	for {
		userOperation := strInput("请输入想要执行的操作:(add/del/mod/query/list/quit) ")
		switch userOperation {
		case "add":
			userAdd(userInput())
			fmt.Println(userList)
		case "del":
			userDel(strInput("请输入要删除的用户ID: "))
		case "mod":
			userMod(strInput("请输入要修改的用户ID: "))
		case "query":
			userQuery(strInput("请输入要查询的信息:(默认扫描所有数据检索) "))
		case "list":
			fmt.Println(userList)
		case "quit":
			break END
		default:
			continue
		}
	}
}
