package services

import (
	"cmdb/models"
	"cmdb/utils"
)

// 查询所有数据并返回特定格式
func APIQueryAllUsers() string {
	data := dataFormat{
		Ok:   "true",
		Msg:  "获取用户数据成功！",
		Data: users,
	}
	return utils.StructJsonMarshal(data)
}

func APIAddUser(u models.User) string {
	// /*
	// 	新增用户案例如下：
	// 		http://127.0.0.1:10086/user/add/?id=4&name=yizuo&sex=true&passwd=123&addr=wo&tel=12334567
	// */
	u.ID = userIdIncrement()
	users = append(users, u)

	data := dataFormat{
		Ok:   "true",
		Msg:  "添加用户数据成功！",
		Data: users,
	}
	return utils.StructJsonMarshal(data)
}

// 根据ID查询特定用户数据
func APIQueryUserID(q string) string {
	k, v := checkUserExists(utils.StrconvAtoi(q))
	if !v {
		data := dataFormat{
			Ok:  "false",
			Msg: "找不到用户ID！",
		}
		return utils.StructJsonMarshal(data)
	}

	data := dataFormat{
		Ok:   "true",
		Msg:  "获取用户数据成功！",
		Data: []models.User{users[k]},
	}
	return utils.StructJsonMarshal(data)
}

// 根据ID删除特定用户的数据
func APIDeleteUser(i string) string {
	k, v := checkUserExists(utils.StrconvAtoi(i))
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

func APIModifyUser(u models.User) string {
	k, v := checkUserExists(u.ID)
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
func userIdIncrement() int {
	i := 0
	for _, v := range users {
		if v.ID > i {
			i = v.ID
		}
	}
	return i + 1
}

// 检查用户是否存在，并返回用户在切片中的位置
func checkUserExists(i int) (int, bool) {
	for k, v := range users {
		if v.ID == i {
			return k, true
		}
	}
	return 0, false
}
