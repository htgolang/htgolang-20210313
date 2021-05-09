package admin

import (
	"cmdb/tools"
	"fmt"
	"strconv"
	"strings"
)

type Users struct {
	Id, Name, Passwd, Addr, Tel string
}

var UserList = make([]*Users, 0)

func NewUser(id, name, passwd, addr, tel string) *Users {
	return &Users{
		Id:     id,
		Name:   name,
		Passwd: passwd,
		Addr:   addr,
		Tel:    tel,
	}
}

// 增加用户
func (u *Users) UserAdd() {
	data := UserInput()
	UserList = append(UserList, &data)
}

// 删除用户
func (u *Users) UserDel() {
	userDelId := tools.StrInput("请输入你要删除的用户ID：")

	userDelPlace, ok := FindIdExists(userDelId)
	if !ok {
		fmt.Println("你输入的用户不存在！")
		return
	}

	fmt.Printf("你要删除的的用户数据为：%v\n", UserList[userDelPlace])
	if !tools.UserInputExecuted("确认要删除用户吗？：(yes/no ; default no)") {
		fmt.Println("输入错误，返回首页。")
		return
	}

	UserList = append(UserList[:userDelPlace], UserList[userDelPlace+1:]...)
	fmt.Println(UserList)

}

// 修改用户
func (u *Users) UserMod() {
	userModId := tools.StrInput("请输入你要修改的用户ID：")

	userModPlace, ok := FindIdExists(userModId)
	if !ok {
		fmt.Println("你输入的用户不存在！")
		return
	}

	fmt.Printf("你要修改的的用户数据为：%v\n", UserList[userModPlace])
	if !tools.UserInputExecuted("确认要修改用户数据吗？：(yes/no ; default no)") {
		fmt.Println("输入错误，返回首页。")
		return
	}

	data := UserInput()
	UserList[userModPlace] = &data
	fmt.Println(UserList)
}

// 检索数据返回
func (u *Users) UserQuery() {
	userNameQuery := tools.StrInput("请输入你要模糊查询的用户名字：")

	userQueryPlace, ok := FindNameExists(userNameQuery)
	if !ok {
		fmt.Println("未查询到你输入的信息。")
		return
	}
	fmt.Printf("查询到的用户数据为：\n%v\n", UserList[userQueryPlace])
}

// 打印当前所有数据
func (u *Users) UserListPrint() {
	for _, v := range UserList {
		fmt.Printf("%#v \n", v)
	}
}

// 查找用户ID是否存在，如果存在返回对应切片位置并返回true，不存在返回false。
func FindIdExists(data string) (int, bool) {
	for k, v := range UserList {
		if v.Id == data {
			return k, true
		}
	}
	return 0, false
}

// 全文检索传入字符串是否存在，如果存在返回对应切片位置并返回true，不存在返回false。
func FindNameExists(data string) (int, bool) {
	for place, mapData := range UserList {
		if strings.Contains(mapData.Name, data) {
			return place, true
		}
	}
	return 0, false
}

// 查找用户ID最大的值+1
func FindMaxIdNum() (maxIdNum int) {
	maxIdNum = 0
	for _, v := range UserList {
		idNum, _ := strconv.Atoi(v.Id)
		if idNum > maxIdNum {
			maxIdNum = idNum + 1

		}
	}
	return maxIdNum
}

// 获取用户输入，返回用户输入的map合集
func UserInput() (data Users) {
	var id, name, passwd, addr, tel string
	id = strconv.Itoa(FindMaxIdNum())
	name = tools.StrInput("请输入用户名称：")
	passwd = tools.StrInput("请输入用户密码：")
	addr = tools.StrInput("请输入用户地址：")
	tel = tools.StrInput("请输入联系方式：")

	data = Users{
		Id:     id,
		Name:   name,
		Passwd: tools.Md5sum(passwd),
		Addr:   addr,
		Tel:    tel,
	}
	return data
}

// 用户登录信息MD5检测
func UserLoginAuth(userData, passwdData string) bool {
	for _, v := range UserList {
		if v.Name == userData && v.Passwd == tools.Md5sum(passwdData) {
			return true
		}
	}
	return false
}

// 用户登录检测
func HandlerUserLoginAuth() bool {
	/*
		循环三次，对比用户输入的用户名密码是否正确
			密码正确返回 true
			密码错误返回 false
	*/
	for i := 0; i < 3; i++ {
		if i != 0 {
			fmt.Println(`密码输入错误，请重新输入密码：`)
		}
		userData := tools.StrInput(`请输入登录用户：`)
		passwdData := tools.StrInput(`请输入密码：`)

		// 对比输入值的md5是否与默认值的md5一致
		if UserLoginAuth(userData, passwdData) {
			return true
		}
	}
	return false
}
