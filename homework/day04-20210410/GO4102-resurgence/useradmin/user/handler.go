package user

import (
	"fmt"
	"strconv"
	. "useradmin/config"
	. "useradmin/utils"
)

func GetUser(userSlice *[]*User) (err error) {
	defer func() {
		if errMsg := recover(); errMsg != nil {
			err = fmt.Errorf("%v", errMsg)
		}
	}()

	if len(*userSlice) == 0 {
		panic("当前没有用户信息")
	}

	for _, m := range *userSlice {
		fmt.Printf("id: %d, name: %s, phone: %s .\n", m.ID, m.Name, m.Phone)
	}
	return
}

func AddUser(userSlice *[]*User) (err error) {
	defer func() {
		if errMsg := recover(); errMsg != nil {
			err = fmt.Errorf("%v", errMsg)
		}
	}()

	name := Input("请输入姓名: ")
	phone := Input("请输入手机号: ")

	//userMap := map[string]string{
	//	"id":    strconv.Itoa(ID),
	//	"name":  name,
	//	"phone": phone,
	//}

	userMap := User{
		ID:    ID,
		Name:  name,
		Phone: phone,
	}

	ID++
	*userSlice = append(*userSlice, &userMap)
	fmt.Printf("%v\n", *userSlice)
	return
}

func DeleteUser(userSlice *[]*User) (err error) {
	defer func() {
		if errMsg := recover(); errMsg != nil {
			err = fmt.Errorf("%v", errMsg)
		}
	}()

	if len(*userSlice) == 0 {
		panic("当前没有用户信息")
	}
	delID := Input("输入要删除的用户id: ")

	// 查找是否存在此id的用户
	for idx, m := range *userSlice {
		if id := strconv.Itoa(m.ID); id == delID {
			// 使用append删除当前索引为 idx 的map
			tmp := *userSlice
			*userSlice = append(tmp[:idx], tmp[idx+1:]...)
			//*userSlice = append(*userSlice[:idx], *userSlice[idx + 1:]...)
			fmt.Println(*userSlice)
			return
		}
	}
	panic(fmt.Sprintf("id为 %s 的用户不存在.", delID))
}

func UpdateUser(userSlice *[]*User) (err error) {
	defer func() {
		if errMsg := recover(); errMsg != nil {
			err = fmt.Errorf("%v", errMsg)
		}
	}()

	if len(*userSlice) == 0 {
		panic("当前没有用户信息")
	}
	var attr, val string

	updateID := Input("输入要修改的用户id: ")

	// 查找是否存在此id的用户
	for _, m := range *userSlice {

		if id := strconv.Itoa(m.ID); id == updateID {

			fmt.Printf("请输入要修改的属性: ")
			fmt.Scan(&attr)
			fmt.Printf("请输入要修改的值: ")
			fmt.Scan(&val)
			fmt.Println("aaaa", attr, val)

			//if _, ok := m[attr]; !ok {
			//	panic("修改的属性不存在")
			//}
			switch attr {
			case "name":
				m.Name = val
			case "phone":
				m.Phone = val
			default:
				fmt.Println("输入不合法")
			}
			fmt.Println(m, *userSlice)
			return
		}
	}
	panic(fmt.Sprintf("id为 %s 的用户不存在.", updateID))
}
