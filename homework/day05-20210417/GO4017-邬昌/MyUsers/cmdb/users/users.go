package users

import (
	"cmdb/models"
	"cmdb/utils"
	"fmt"
	"strconv"
)

var id int
var name, addr, tel string

//var user = models.Us

var user = models.UserList{make(map[string]*models.User)}

func getId() int {

	for _, value := range user.Userlist {
		i, _ := strconv.Atoi(value.Id)
		fmt.Println("i的值为：", i)

		if i > id {
			id = i
		}
	}
	return id + 1
}

//添加用户
func AddUser() {

	fmt.Println("---执行添加用户操作---")
	name := utils.Input("请输入用户名：")
	addr := utils.Input("请输入地址：")
	tel := utils.Input("请输入电话号码：")

	//创建用户，调用User的构造函数
	newUser := models.NewUser(strconv.Itoa(getId()), name, addr, tel)
	//追加到map中
	user.Userlist[strconv.Itoa(getId())] = newUser

}

//删除用户
func DeletUser() {

	fmt.Println("---执行删除操作---")
	delid := utils.Input("请输入要删除的用户id：")

	_, ok := IfUserExist(delid)
	if !ok {
		fmt.Println("需要查找的用户不存在！！！")
		return
	}

	str := utils.Input("请输入y/yes/Y/YES确认删除:")

	if str == "y" || str == "yes" || str == "Y" || str == "YES" {

		delete(user.Userlist, delid)

		fmt.Println("用户", user.Userlist[delid], "被删除了！！！")

		fmt.Println(user)

	} else {
		return
	}

}

//查询用户
func QueryUsers() {

	fmt.Println("---执行查询用户操作---")
	querystr := utils.Input("请输入要查询的用户名:")
	//判断map中是否存在输入的元素
	for _, value := range user.Userlist {
		if value.Name == querystr {

			fmt.Println(value.Id, value.Name, value.Addr, value.Tel)
		} else {

			fmt.Println("查无此人")
		}
	}
}

//修改用户
func ModifyUsers() {

	fmt.Println("---执行修改用户操作---")
	modifyid := utils.Input("请输入要修改的用户id：")

	us, ok := IfUserExist(modifyid)
	if !ok {

		fmt.Println("查无此人")
		return
	}

	str := utils.Input("请输入y/yes/Y/YES确认修改:")

	if str == "y" || str == "yes" || str == "Y" || str == "YES" {
		name := utils.Input("请输入用户名：")

		addr := utils.Input("请输入地址：")

		tel := utils.Input("请输入电话号码；")

		us.Name = name
		us.Addr = addr
		us.Tel = tel
		newUser := models.NewUser(strconv.Itoa(getId()), name, addr, tel)
		user.Userlist[strconv.Itoa(getId())] = newUser

	} else {

		return
	}

}

func IfUserExist(userid string) (models.User, bool) {

	for _, value := range user.Userlist {
		if userid == value.Id {

			return *value, true
		}
	}
	return models.User{}, false

}
