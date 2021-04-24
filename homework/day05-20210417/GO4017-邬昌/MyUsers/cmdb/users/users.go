package users

import (
	"cmdb/models"
	"fmt"
	"strconv"
)

var id int
var name, addr,tel string

var user = models.Us


func getId() int  {

	for _,value :=range user{
		i,_ := strconv.Atoi(value.Id)
		fmt.Println("i的值为：",i)

		if i > id{
			id =i
		}}
	return id+1
}


func Input(str string) string  {

	var text string
	fmt.Println(str)
	fmt.Scan(&text)
	return text

}

//添加用户
func AddUser() {

	fmt.Println("---执行添加用户操作---")
	name :=Input("请输入用户名：")
	addr :=Input("请输入地址：")
	tel :=Input("请输入电话号码：")

	//创建用户，调用User的构造函数
	newUser :=models.NewUser(strconv.Itoa(getId()),name,addr,tel)
	//追加到map中
	user[strconv.Itoa(getId())]=newUser

}

//删除用户
func DeletUser() {

	fmt.Println("---执行删除操作---")
	delid :=Input("请输入要删除的用户id：")

	_,ok :=IfUserExist(delid)
	if !ok{
		fmt.Println("需要查找的用户不存在！！！")
		return
	}

	str :=Input("请输入y/yes/Y/YES确认删除:")

	if (str =="y" || str =="yes" || str =="Y" || str =="YES" ){

		delete(user, delid)

		fmt.Println("用户",user[delid],"被删除了！！！")

		fmt.Println(user)

	}else {
		return
	}

}

//查询用户
func QueryUsers(){

	fmt.Println("---执行查询用户操作---")
	querystr :=Input("请输入要查询的用户名:")
	//判断map中是否存在输入的元素
	for _,value :=range user{
		if value.Name ==querystr{

			fmt.Println(value.Id,value.Name,value.Addr,value.Tel)
		}else {

			fmt.Println("查无此人")
		}
	}
}

//修改用户
func ModifyUsers() {

	fmt.Println("---执行修改用户操作---")
	modifyid :=Input("请输入要修改的用户id：")

	us,ok :=IfUserExist(modifyid)
	if !ok {

		fmt.Println("查无此人")
		return
	}

	str :=Input("请输入y/yes/Y/YES确认修改:")

	if (str =="y" || str =="yes" || str =="Y" || str =="YES" ){
	    	name :=Input("请输入用户名：")

		    addr :=Input("请输入地址：")

			tel :=Input("请输入电话号码；")

			us.Name = name
			us.Addr = addr
			us.Tel  = tel
		    newUser :=models.NewUser(strconv.Itoa(getId()),name,addr,tel)
		    user[strconv.Itoa(getId())] =newUser

	} else {

		return
	}

}

func IfUserExist(userid string)(models.User,bool){

	for _,value :=range user{
		if userid == value.Id{

			return *value,true
		}
	}
	return models.User{},false

}