package handlers

import (
	"cmdb/models"
	"cmdb/services"
	"cmdb/utils"
	"fmt"
	"net/http"
)

// 查询所有用户控制器控制器
func APIQueryAllUserHandler(w http.ResponseWriter, r *http.Request) {
	data := services.APIQueryAllUsers()
	fmt.Fprint(w, data)
}

// 查询控制器
func APIQueryUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.FormValue("id") == "" {
		data := dataFormat{
			Ok:  "false",
			Msg: "用户ID为空！",
		}
		fmt.Fprint(w, utils.StructJsonMarshal(data))
		return
	}
	data := services.APIQueryUserID(r.FormValue("id"))
	fmt.Fprint(w, data)
}

// 添加控制器
func APIAddUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form) < 5 {
		data := utils.StructJsonMarshal(dataFormat{
			Ok:  "false",
			Msg: "用户数据缺失，需要 name,sex,passwd,addr,tel.",
		})
		fmt.Fprint(w, data)
		return
	}
	user := models.User{
		Name:   r.FormValue("name"),
		Sex:    utils.TextSex(r.FormValue("sex")),
		Passwd: utils.Md5sum(r.FormValue("passwd")),
		Addr:   r.FormValue("addr"),
		Tel:    r.FormValue("tel"),
	}
	data := services.APIAddUser(user)
	fmt.Fprint(w, data)
}

// 删除控制器
func APIDeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.FormValue("id") == "" {
		data := dataFormat{
			Ok:  "false",
			Msg: "用户ID为空！",
		}
		fmt.Fprint(w, utils.StructJsonMarshal(data))
		return
	}
	data := services.APIDeleteUser(r.FormValue("id"))
	fmt.Fprint(w, data)
}

// 变更控制器
func APIModifyUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form) < 6 {
		data := utils.StructJsonMarshal(dataFormat{
			Ok:  "false",
			Msg: "用户数据缺失，需要 id,name,sex,passwd,addr,tel.",
		})
		fmt.Fprint(w, data)
		return
	}
	user := models.User{
		ID:     utils.StrconvAtoi(r.FormValue("id")),
		Name:   r.FormValue("name"),
		Sex:    utils.TextSex(r.FormValue("sex")),
		Passwd: utils.Md5sum(r.FormValue("passwd")),
		Addr:   r.FormValue("addr"),
		Tel:    r.FormValue("tel"),
	}
	data := services.APIModifyUser(user)
	fmt.Fprint(w, data)
}
