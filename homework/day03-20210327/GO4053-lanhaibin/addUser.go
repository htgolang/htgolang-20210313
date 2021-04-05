package main

import (
	"fmt"
	"strconv"
)
var users = []map[string]string{
	{
		"id": "1",
		"name": "foo",
		"address": "beijing",
	},
	{
		"id": "2",
		"name": "bar",
		"address": "shanghai",
	},
	{
		"id": "3",
		"name": "foobar",
		"address": "guangzhou",
	},
}

func addUser() {
	user := map[string]string{"id":"","name":"","address":""}
	var id, name, address string
	id = getId()
	fmt.Println("开始添加用户:")
	fmt.Print("请输入用户名:")
	fmt.Scanf("%s", &name)
	fmt.Print("请输入地址:")
	fmt.Scanf("%s", &address)
	user["id"] = id
    user["name"] = name
    user["address"] = address
	users = append(users, user)
}

func getId() string {
	var id, maxid int64
	for _, v := range users {
		id, _ = strconv.ParseInt(v["id"],10, 64)
		if id > maxid {
			maxid = id
		}
	}
	return strconv.FormatInt(id, 10)
}

func main() {
	addUser()
	fmt.Println(users)
}