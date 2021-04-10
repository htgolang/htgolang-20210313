package main

import (
	"fmt"
	"strconv"
	"strings"
	// "unicode"
)

//1 用户数据结构; 输入相关序号执行相关功能（函数块实现）
//2 用户添加
/* id
username
addr
phone
*/	
func userAdd(id int, users map[string]map[string]string) map[string]map[string]string {
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

	users[strconv.Itoa(id)] =  map[string]string{
		"username": username,
		"email": addr,
		// "password": password,
		"phone": phone,
	}

	fmt.Println(users)
	return users
}

//3 用户查询
func query(users map[string]map[string]string) {
	var q string
	fmt.Println("请输入查询:")
	fmt.Scan(&q)
	for k, v := range users{
		if len(q) == 0{
			fmt.Println(users)		
		} else if strings.Contains(v["username"], q) || strings.Contains(v["addr"], q) ||  strings.Contains(v["phone"],q){
			fmt.Printf("%5s %5s %5s %5s\n", k, v["username"], v["email"], v["phone"])
		}
	}
}

//4 用户删除
func delUser(users map[string]map[string]string){
	var delstr string
	var dete string
	fmt.Println("请输入要删除的用户:")
	fmt.Scan(&delstr)
	for k, v := range users{
		if delstr == v["username"]{
			fmt.Printf("是否要删除%v用户(y/yes/Y/YES):")
			fmt.Scan(&dete)
			if dete == "y"|| dete == "yes"|| dete == "Y"|| dete == "YES"{
				delete(users, k)
			} else {
				return
			}
		}
	}
	fmt.Println(users)
}

//5 用户信息修改
func modifUser(users map[string]map[string]string){
	var modstr, orgistr, addrs, phones string
	fmt.Println("请输入要修改的用户名:")
	fmt.Scan(&orgistr)
	for _,v := range users{
		if orgistr == v["username"]{
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

func main(){
	users := map[string]map[string]string{}
	var n string
	id := 0

	fmt.Println("xxx用户系统")
	for {
		fmt.Println(`
1. 用户添加
2. 用户查询
3. 用户删除
4. 用户修改
5. 退出`)
		fmt.Println("请输入编号:")
		fmt.Scan(&n)
		switch {
		case n == "1":
			id++
			userAdd(id, users)
		case n == "2":
			query(users)
		case n == "3":
			delUser(users)
		case n == "4":
			modifUser(users)
		case n == "5":
			return
		}
	}
}






