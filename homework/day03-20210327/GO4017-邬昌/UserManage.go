package main

import (
	"fmt"
	"strconv"
)


var id int
var name, addr,tel string

var users = []map[string]string{
	{
		"id"  : "1",
		"name": "小明",
		"addr": "北京",
		"tel" : "18079220200",
	},
}
func getId() int  {

	for _,value :=range users{
		i,_ := strconv.Atoi(value["id"])
		fmt.Println("i的值为：",i)
		if i > id{
			id =i
		}}
	return id+1
}
//添加用户

func input(str string) string  {

	var text string
	fmt.Println(str)
	fmt.Scan(&text)
	return text

}
func addUser()([]map[string]string)  {

	name :=input("请输入用户名：")
	addr :=input("请输入地址：")
	tel :=input("请输入电话号码：")

	for _,value :=range users{
		if value["name"] ==name{

			fmt.Println("用户%s已存在",name)
			return users
		}

	}

	users =append(users, map[string]string{
			"id": strconv.Itoa(getId()),
			"name": name,
			"addr": addr,
			"tel":  tel,
	})

	return users
}

//删除用户
func deletUser()[]map[string]string {

	delid :=input("请输入要删除的用户id：")

	//遍历users查看是否存在key
	for _,value :=range users{
		fmt.Println("value",value["id"])
		if delid ==value["id"]{
			fmt.Println("用户信息为：",users)

			str :=input("请输入y/yes/Y/YES确认删除:")

			if (str =="y" || str =="yes" || str =="Y" || str =="YES" ){

				i,_:=strconv.Atoi(delid)
				users = append(users[:i-1], users[i:]...)
				fmt.Println("用户",delid,"被删除了！！！")
				return users
			}
		}else {
			    fmt.Println("查无此人")
		}
	}

	return nil
}

//查询用户
func queryUsers()[]map[string]string{


	querystr :=input("请输入要查询的用户名:")
	//判断map中是否存在输入的元素
	for _,value :=range users{
		if value["name"] ==querystr{

			return users
		}else {

			fmt.Println("查无此人")
		}
	}

	return nil
}

//修改用户
func modifyUsers() []map[string]string{

	modifyid :=input("请输入要修改的用户id：")

	for _,value :=range users{
		if modifyid ==value["id"]{
			fmt.Println("用户信息为：",users)
			str :=input("请输入y/yes/Y/YES确认修改:")

			if (str =="y" || str =="yes" || str =="Y" || str =="YES" ){
				name :=input("请输入用户名：")

			    addr :=input("请输入地址：")

				tel :=input("请输入电话号码；")

				users =append(users, map[string]string{
					"id": modifyid,
					"name": name,
					"addr": addr,
					"tel":  tel,
				})

				return users

			}else{
				break
			}

		}else {

			fmt.Println("查无此人")

		}
	}

	return nil
}

func main() {

	var do int

	/*
		执行操作：
		1、添加用户  2、删除用户
		3、查询用户  4、修改用户
	*/
	for {
		fmt.Println("请输入要执行的操作,1:添加用户,2:删除用户,3:查询用户,4:修改用户,其他按键退出")
		fmt.Scan(&do)

		if do == 1 {

			fmt.Println(addUser())
		} else if do == 2 {

			fmt.Println(deletUser())
		} else if do == 3 {

			fmt.Println(queryUsers())
		} else if do == 4 {

			fmt.Println(modifyUsers())
		} else if do >4{

			fmt.Println("退出游戏！！！")
			return
		}
	}

}
