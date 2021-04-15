package user

import (
	"fmt"
	"strconv"
	"strings"
)


var users = []map[string]string {
	{
		"id":"1",
		"name":"arun",
		"addr": "Beijing",
		"tel": "xxx123123123",
	},
	{
		"id" : "2",
		"name": "ujianfeng",
		"addr":"Beijing",
		"tel":"152xxxxxxx",
	},
}

func genId() int{
	id := 0
	for _,user := range users {
		i, _ := strconv.Atoi(user["id"])
		if i > id {
			id = i
		}
	}
	return id + 1
}

func input(prompt string) string {
	fmt.Println(prompt)
	var text string
	fmt.Scan(&text)
	return text
}

//添加用户
func Add() {
	name := input("请输入用户名:")
	addr := input("请输入地址:")
	tel := input("请输入电话:")

	//验证用户名不能重复

	for k,_ := range users {
		if name == users[k]["name"] {
			fmt.Errorf("用户名%s已存在",name)
		}
	}
	// if name == "kk" {
	// 	return fmt.Errorf("用户名%s已存在",name)
	// }
	users = append(users,map[string]string{
		"id": strconv.Itoa(genId()),
		"name" : name,
		"addr" : addr,
		"tel" : tel,
	})
	//return nil 
}

//删除用户
func Delete() {
	id := 0
	fmt.Println("请输入要删除的用户ID:")
	fmt.Scan(&id)
	for i, _ := range users{
		if users[i]["id"] == strconv.Itoa(id) {
			// fmt.Println("你想删除用户:",users[k])
			fmt.Println(strings.Repeat("*",10))
			for k,v := range users[i] {
				fmt.Println(k,v)
			}
			fmt.Println(strings.Repeat("*",10))
			
		var confirm string
		fmt.Println("你确认要删除这个用户？(y/yes/Y/YES)")
		fmt.Scan(&confirm)
		if confirm == "y" || confirm =="yes" || confirm =="Y" || confirm =="YES" {
			// if i != len(users){
			// 	users = append(users[:i],users[i+1:]...)
			// } else {
			// 	users = users[:len(users)-1]
			// }
			copy(users[i:],users[i+1:])
			users = users[:len(users)-1]
			//fmt.Println(users)
			//return users
			
			//users = users[:len(users)-1]
			} 
	}
}
	//return users
}

//查询用户

func Query(){
	var q string
	fmt.Println("请输入查询:(姓名/地址/电话)")
	fmt.Scan(&q)

	for _, v := range users {
		if len(q) == 0 {
			fmt.Println(users)
		} else if strings.Contains(v["name"],q) || strings.Contains(v["addr"],q) || strings.Contains(v["tel"],q) {
			fmt.Printf("%s %s %s %s\n",v["id"], v["name"],v["addr"],v["tel"])
		}
	}
}


//修改用户信息
func Modify() {
	var nname,orginame,addrs,phones string
	fmt.Println("请输入要修改用户名:")
	fmt.Scan(&orginame)
	for k,_ := range users {
		if orginame == users[k]["name"] {
			fmt.Println("请输入新的用户名:")
			fmt.Scan(&nname)
			fmt.Println("请输入新地址:")
			fmt.Scan(&addrs)
			fmt.Println("请输入新电话:")
			fmt.Scan(&phones)
			users[k]["name"] = nname
			users[k]["addr"] = addrs
			users[k]["tel"] = phones
		} 
	}
	fmt.Println(users)
}
