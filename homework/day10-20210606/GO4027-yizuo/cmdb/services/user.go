package services

import (
	"cmdb/models"
	"cmdb/utils"
)

// 查询所有数据并返回特定格式
func QueryAllUsers() []models.User {
	return users
}

// 添加用户
func AddUser(u models.User) {
	u.ID = UserIdIncrement()
	users = append(users, u)
}

// 根据ID查询特定用户数据
func QueryUserID(q string) (models.User, bool) {
	k, v := CheckUserExists(utils.StrconvAtoi(q))
	if !v {
		return models.User{}, false
	}
	return users[k], true
}

// 根据ID删除特定用户的数据
func DeleteUser(i string) string {
	k, v := CheckUserExists(utils.StrconvAtoi(i))
	if v {
		users = append(users[:k], users[k+1:]...)
		data := dataFormat{
			Ok:   "true",
			Msg:  "删除用户数据成功！",
			Data: users,
		}
		return utils.StructJsonMarshal(data)
	}

	data := dataFormat{
		Ok:  "false",
		Msg: "找不到用户ID！",
	}
	return utils.StructJsonMarshal(data)
}

// 变更用户数据
func ModifyUser(u models.User) string {
	k, v := CheckUserExists(u.ID)
	if v {
		users[k] = u
		data := dataFormat{
			Ok:   "true",
			Msg:  "用户数据修改成功！",
			Data: users,
		}
		return utils.StructJsonMarshal(data)
	}

	data := dataFormat{
		Ok:  "false",
		Msg: "找不到用户ID！",
	}
	return utils.StructJsonMarshal(data)
}

// 返回一个不会重复的用户ID
func UserIdIncrement() int {
	i := 0
	for _, v := range users {
		if v.ID > i {
			i = v.ID
		}
	}
	return i + 1
}

// 检查用户是否存在，并返回用户在切片中的位置
func CheckUserExists(i int) (int, bool) {
	for k, v := range users {
		if v.ID == i {
			return k, true
		}
	}
	return 0, false
}

// 检查用户密码是否正确
func CheckUserPasswordLogin(u, p string) *models.User {
	for _, v := range users {
		if v.Name == u && v.Passwd == utils.Md5sum(p) {
			return &v
		}
	}
	return nil
}
