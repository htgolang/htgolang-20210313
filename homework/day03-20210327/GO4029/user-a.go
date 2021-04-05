package main

import "fmt"

func main() {
	var usersSlice = make([]map[string]string, 0)
	var users = make(map[string]string)
	var userName string
	var address string
	var phone string
	var con string
	for {
		fmt.Println("请输入用户名：")
		fmt.Scan(&userName)
		fmt.Println("请输入地址：")
		fmt.Scan(&address)
		fmt.Println("请输入电话号码：")
		fmt.Scan(&phone)
		userId := len(usersSlice) + 1
		users["id"] = fmt.Sprintf("%d", userId)
		users["name"] = userName
		users["addr"] = address
		users["tel"] = phone
		usersSlice = append(usersSlice, users)
		fmt.Printf("当前列表：%v", usersSlice)
		fmt.Println("是否继续添加，Y/N:")
		fmt.Scan(&con)
		if con == "Y" {
			continue
		} else {
			break
		}
	}
}
