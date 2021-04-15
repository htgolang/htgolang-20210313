package main
// 为什么有index out of range这个错误
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
	{
		"id" : "3",
		"name" : "John",
		"addr" : "Shanghai",
		"tel" : "132xxxxx",
	},
}


//删除用户
func deluser() []map[string]string{
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
			return users
			
			//users = users[:len(users)-1]
			} 
	}
}
	return users
}

func main() {
	
	fmt.Println(deluser())
	

}