package user

import (
	"fmt"
	"mgr/model/user"
	"strings"
)

func FindUser(users []*user.User)  {
	fmt.Printf("请输入要模糊查询的字符串:")
	var searchStr string
	flag:=false
	fmt.Scan(&searchStr)
	//users = Create("gob").Read()
	users = Create("json").Read()
	//users = Create("csv").Read()
	for i := 0; i < len(users); i++ {
		if strings.Contains(users[i].Name,searchStr)||strings.Contains(users[i].Addr,searchStr)||strings.Contains(users[i].Tel,searchStr){
			fmt.Printf("id:%v,name:%v,tel:%v,addr:%v;\n",users[i].Id,users[i].Name,users[i].Tel,users[i].Addr)
			flag = true
		}
	}
	if !flag{
		fmt.Println("您搜索的用户不存在!")
	}
}

