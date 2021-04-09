package main

import (
	"fmt"
	"strconv"
)

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
//添加用户
func addUser()(map[string]map[string]string)  {

	var name, addr,tel string
	var users = make(map[string]map[string]string)
	var id int
	fmt.Println("请输入用户名：")
	fmt.Scan(&name)

	fmt.Println("请输入家庭地址：")
	fmt.Scan(&addr)

	fmt.Println("请输入电话号码：")
	fmt.Scan(&tel)
	//定义id自增
	id++
	users[strconv.Itoa(id)] = map[string]string{
		"name": name,
		"addr": addr,
		"tel":  tel,
	}

	return users
}

//删除用户
func deletUser()map[string]map[string]string {
	var users map[string]map[string]string
	var delid int
	var str string
	fmt.Println("请输入要删除的用户id：")
	fmt.Scan(&delid)

	//遍历users查看是否存在key
	for key,_ :=range users{
		if strconv.Itoa(delid) ==key{
			fmt.Println("用户信息为：",users)
			fmt.Println("请输入y/yes/Y/YES确认删除:")
			fmt.Scan(&str)
			if (str =="y" || str =="yes" || str =="Y" || str =="YES" ){

				delete(users, strconv.Itoa(delid))
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
func queryUsers()map[string]map[string]string{

	var users map[string]map[string]string
	var querystr string
	fmt.Println("请输入要查询的用户名或者地址或者电话")
	fmt.Scan(&querystr)
	//判断map中是否存在输入的元素
	for _,value :=range users{
		for _,val :=range value{
			if val ==querystr{
				return users
			}

		}
	}

	return nil
}

//修改用户
func modifyUsers() map[string]map[string]string{

	var users map[string]map[string]string
	var modifyid int
	var str string
	var name,addr,tel string
	fmt.Println("请输入要删除的用户id：")
	fmt.Scan(&modifyid)
	for key,_ :=range users{
		if strconv.Itoa(modifyid) ==key{
			fmt.Println("用户信息为：",users)
			fmt.Println("请输入y/yes/Y/YES确认修改:")
			fmt.Scan(&str)
			if (str =="y" || str =="yes" || str =="Y" || str =="YES" ){
				fmt.Println("请输入用户名：")
				fmt.Scan(&name)

				fmt.Println("请输入地址：")
				fmt.Scan(&addr)

				fmt.Println("请输入电话号码；")
				fmt.Scan(&tel)
				users[strconv.Itoa(modifyid)] = map[string]string{
					"name": name,
					"addr": addr,
					"tel":  tel,
				}

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