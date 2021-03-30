package main

import (
"fmt"
"strings"
)

func findUser(users []map[string]string)  {
	fmt.Printf("请输入要模糊查询的字符串:")
	var searchStr string
	flag:=false
	fmt.Scan(&searchStr)
	for i := 0; i < len(users); i++ {
		if strings.Contains(users[i]["name"],searchStr)||strings.Contains(users[i]["addr"],searchStr)||strings.Contains(users[i]["tel"],searchStr){
			fmt.Printf("id:%v,name:%v,tel:%v,addr:%v;\n",users[i]["id"],users[i]["name"],users[i]["tel"],users[i]["addr"])
			flag = true
		}
	}
	if !flag{
		fmt.Println("您搜索的用户不存在!")
	}
}

