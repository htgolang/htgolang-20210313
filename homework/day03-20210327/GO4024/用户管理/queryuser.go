package main

import (
	"fmt"
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
	{
		"id" : "3",
		"name" : "John",
		"addr" : "Shanghai",
		"tel" : "132xxxxx",
	},
}
//用户查询

func queryuser(){
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

func main() {

	queryuser()
}