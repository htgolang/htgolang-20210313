package user

import (
	"fmt"
	"strconv"
	"strings"
)

//2 用户添加
/* id
username
addr
phone
*/
// var users map[string]map[string]string = map[string]map[string]string{}
// var n string
// var id int = 1
func UserAdd(id int, users map[string]map[string]string) map[string]map[string]string {
	var username string
	// var password string
	var phone string
	var addr string
	fmt.Println("请输入用户名:")
	fmt.Scan(&username)
	fmt.Println("请输入地址:")
	fmt.Scan(&addr)
	// fmt.Println("请输入密码:")
	// fmt.Scan(&password)
	fmt.Println("请输入联系方式:")
	fmt.Scan(&phone)

	users[strconv.Itoa(id)] = map[string]string{
		"username": username,
		"email":    addr,
		// "password": password,
		"phone": phone,
	}

	fmt.Println(users)
	id++
	return users
}

//3 用户查询
func Query(users map[string]map[string]string) {
	var q string
	fmt.Println("请输入查询用户:")
	fmt.Scan(&q)
	for k, v := range users {
		if len(q) == 0 {
			fmt.Println(users)
		} else if strings.Contains(v["username"], q) || strings.Contains(v["addr"], q) || strings.Contains(v["phone"], q) {
			fmt.Printf("id:%2s  username:%2s  email:%2s  phone:%2s\n", k, v["username"], v["email"], v["phone"])
		}
	}
}

//4 用户删除
func DelUser(users map[string]map[string]string) {
	var delstr string
	var dete string
	fmt.Println("请输入要删除的用户:")
	fmt.Scan(&delstr)
	for k, v := range users {
		if delstr == v["username"] {
			fmt.Printf("是否要删除%v用户(y/yes/Y/YES):")
			fmt.Scan(&dete)
			if dete == "y" || dete == "yes" || dete == "Y" || dete == "YES" {
				delete(users, k)
			} else {
				return
			}
		}
	}
	fmt.Println(users)
}

//5 用户信息修改
func ModifUser(users map[string]map[string]string) {
	var modstr, orgistr, addrs, phones string
	fmt.Println("请输入要修改的用户名:")
	fmt.Scan(&orgistr)
	for _, v := range users {
		if orgistr == v["username"] {
			fmt.Println("请输入新的用户名:")
			fmt.Scan(&modstr)
			fmt.Println("请输入新地址:")
			fmt.Scan(&modstr)
			fmt.Println("请输入新联系电话:")
			fmt.Scan(&modstr)
			v["username"] = modstr
			v["addr"] = addrs
			v["phone"] = phones
		}
	}
	fmt.Println(users)
}
