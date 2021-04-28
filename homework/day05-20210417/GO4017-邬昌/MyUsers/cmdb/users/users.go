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

		if i > id {
			id = i
		}
	}
	return id + 1
}

//添加用户
func AddUser() {

	fmt.Println("---执行添加用户操作---")

	//创建用户，调用User的构造函数

	if err := CheckUser(&user); err == nil {

		newUser := models.NewUser(strconv.Itoa(getId()), name, addr, tel)

		user.Userlist[strconv.Itoa(getId())] = newUser
		fmt.Println("[+]用户添加成功！！！")

	} else {
		fmt.Println(err)
	}

	//追加到map中

	//fmt.Println(models.UserList{})

}

func CheckUser(Us *models.UserList) error {

	name = utils.Input("请输入用户名：")
	addr = utils.Input("请输入地址：")
	tel = utils.Input("请输入电话号码：")

	for _, value := range Us.Userlist {

		if name == "" {

			return fmt.Errorf("输入的用户名为空！！！")
		} else if value.Name == name {

			return fmt.Errorf("用户名重复！！！")
		}

	}
	return nil

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
	queryselect := utils.Input("请输入要执行的查询操作:1、查询全部  2、按条件索引")
	//判断map中是否存在输入的元素
	for _, value := range user.Userlist {
		if queryselect == "1" {

			fmt.Printf("用户id: %v 用户名: %v 地址: %v 电话: %v\n", value.Id, value.Name, value.Addr, value.Tel)
		} else if queryselect == "2" {

			querystr := utils.Input("请输入要查询的姓名：")
			if querystr == value.Name {

				fmt.Printf("用户id: %v 用户名: %v 地址: %v 电话: %v\n", value.Id, value.Name, value.Addr, value.Tel)
				return
			}
		} else {

			break
		}
	}

}

//修改用户
func ModifyUsers() {

	fmt.Println("---执行修改用户操作---")
	modifyId := utils.Input("请输入要修改的用户id：")

	us, ok := IfUserExist(modifyId)
	if !ok {

		fmt.Println("查无此人")
		return
	}

	fmt.Println("将修改的用户信息：", us)
	str := utils.Input("请输入y/yes确认修改:")

	if str == "y" || str == "yes" {

		if err := CheckModifyUser(&user); err == nil {

			us = *models.NewUser(modifyId, name, addr, tel)
			user.Userlist[modifyId] = &us
			fmt.Println("修改用户成功！！！")

		}else {

			fmt.Println(err)
		}

	} else {

		return
	}

}

func CheckModifyUser(Us *models.UserList) error {

	name = utils.Input("请输入用户名：")
	addr = utils.Input("请输入地址：")
	tel = utils.Input("请输入电话号码：")

	for _, value := range Us.Userlist {

		if value.Name == name && value.Id == strconv.Itoa(getId()) {

			return fmt.Errorf("用户名重复！！！")
		}

	}
	return nil

}

func IfUserExist(userid string) (models.User, bool) {

	for _, value := range user.Userlist {
		if userid == value.Id {

			return *value, true
		}
	}
	return models.User{}, false

}
