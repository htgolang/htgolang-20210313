package main
// 想按函数功能区分开像adduser一样，但是没有思路
//执行结果有nil
import (
	"fmt"
	"strconv"
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

// func input(prompt int) int {
// 	fmt.Println(prompt)
// 	var id int
// 	fmt.Scan(&id)
// 	return id

func deluser() error {

	id := 0
	fmt.Println("请输入要删除的用户ID:")
	fmt.Scan(&id)
	
	for i, _ := range users{
		if users[i]["id"] == strconv.Itoa(id) {
			// fmt.Println("你想删除用户:",users[k])
			fmt.Println("===========用户信息如下：=============")
			for k,v := range users[i] {
				fmt.Println(k,v)
			}
			fmt.Println("===================================")
			
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
			fmt.Println(users)
			//return users
			
			//users = users[:len(users)-1]
			
			return nil
			break
			}
	} 
}
	return fmt.Errorf("用户%d不存在",id)
	
}

func main() {
	
	fmt.Println(deluser())
	

}