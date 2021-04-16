package users

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

func Input(str string) string  {

	var text string
	fmt.Println(str)
	fmt.Scan(&text)
	return text

}
func AddUser()([]map[string]string)  {

	fmt.Println("---执行添加用户操作---")
	name :=Input("请输入用户名：")
	addr :=Input("请输入地址：")
	tel :=Input("请输入电话号码：")

	for _,value :=range users{
		if value["name"] ==name{

			fmt.Println("用户%s已存在，请重新创建！！！",name)
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
func DeletUser()[]map[string]string {

	fmt.Println("---执行删除操作---")
	delid :=Input("请输入要删除的用户id：")

	//遍历users查看是否存在key
	for _,value :=range users{

		fmt.Println("value",value["id"])

		if delid ==value["id"]{
			fmt.Println("用户信息为：",users)

			str :=Input("请输入y/yes/Y/YES确认删除:")

			if (str =="y" || str =="yes" || str =="Y" || str =="YES" ){

				i,_:=strconv.Atoi(delid)
				users = append(users[:i-1], users[i:]...)
				fmt.Println("用户",value["name"],"被删除了！！！")
				return users
			}
		}else {
			fmt.Println("查无此人")
		}
	}

	return nil
}

//查询用户
func QueryUsers()[]map[string]string{

	fmt.Println("---执行查询用户操作---")
	querystr :=Input("请输入要查询的用户名:")
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
func ModifyUsers() []map[string]string{

	fmt.Println("---执行修改用户操作---")
	modifyid :=Input("请输入要修改的用户id：")

	for _,value :=range users{
		if modifyid ==value["id"]{
			fmt.Println("用户信息为：",users)
			str :=Input("请输入y/yes/Y/YES确认修改:")

			if (str =="y" || str =="yes" || str =="Y" || str =="YES" ){
				name :=Input("请输入用户名：")

				addr :=Input("请输入地址：")

				tel :=Input("请输入电话号码；")

				value["name"] = name
				value["addr"] = addr
				value["tel"]  = tel

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

func Prom()  {
	fmt.Println(`
    1、添加用户
    2、删除用户
    3、查询用户
    4、修改用户
    5、其他按键退出
     `)
}
