package services

import (
	"cmdb/models"
	"cmdb/utils"
	"encoding/json"
)

type dataFormat struct {
	Ok   string      `json:"ok,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

var users = []models.User{
	models.User{1, "yizuo", utils.Md5sum("yizuo"), true, "Wuhan", "66666666"},
	models.User{2, "saber", utils.Md5sum("saber"), false, "Wuhan", "66666666"},
	models.User{3, "leimu", utils.Md5sum("leimu"), false, "Shanghai", "88888888"},
}

// 查询所有数据并返回特定格式
func QueryAllUsers() string {
	data := dataFormat{
		Ok:   "true",
		Msg:  "获取用户数据成功！",
		Data: users,
	}
	return StructJsonMarshal(data)
}

func AddUser(u models.User) string {
	// /*
	// 	新增用户案例如下：
	// 		http://127.0.0.1:10086/user/add/?id=4&name=yizuo&sex=true&passwd=123&addr=wo&tel=12334567
	// */
	u.ID = UserIdIncrement()
	users = append(users, u)

	data := dataFormat{
		Ok:   "true",
		Msg:  "添加用户数据成功！",
		Data: users,
	}
	return StructJsonMarshal(data)
}

// 根据ID查询特定用户数据
func QueryUserID(q string) string {
	k, v := CheckUserExists(utils.StrconvAtoi(q))
	if !v {
		data := dataFormat{
			Ok:  "false",
			Msg: "找不到用户ID！",
		}
		return StructJsonMarshal(data)
	}

	data := dataFormat{
		Ok:   "true",
		Msg:  "获取用户数据成功！",
		Data: users[k],
	}
	return StructJsonMarshal(data)
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
		return StructJsonMarshal(data)
	}

	data := dataFormat{
		Ok:  "false",
		Msg: "找不到用户ID！",
	}
	return StructJsonMarshal(data)
}

func ModifyUser(u models.User) string {
	k, v := CheckUserExists(u.ID)
	if v {
		users[k] = u
		data := dataFormat{
			Ok:   "true",
			Msg:  "用户数据修改成功！",
			Data: users,
		}
		return StructJsonMarshal(data)
	}

	data := dataFormat{
		Ok:  "false",
		Msg: "找不到用户ID！",
	}
	return StructJsonMarshal(data)
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

// 结构体转换为json格式的字符串并返回
func StructJsonMarshal(i interface{}) string {
	data, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(data)
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
