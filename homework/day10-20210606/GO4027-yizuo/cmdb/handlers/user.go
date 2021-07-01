package handlers

import (
	"cmdb/models"
	"cmdb/services"
	"cmdb/utils"
	"cmdb/web"
	"fmt"
	"html/template"
	"net/http"
)

// 用户登录控制器
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	session := web.LoadSession(w, r)
	if ok := session.Get("uid"); ok != nil {
		fmt.Println("123", session.Get("uid"))
		fmt.Println(session)
		fmt.Println(ok)
		http.Redirect(w, r, "/user", http.StatusFound)
	}

	var name, password, Msg string
	if r.Method == "POST" {
		name = r.FormValue("name")
		password = r.FormValue("password")
		if user := services.CheckUserPasswordLogin(name, password); user != nil {
			session.Set("uid", user.ID)
			if err := web.DumpSession(w, r, session); err != nil {
				fmt.Println(err)
			}
			http.Redirect(w, r, "/user", http.StatusFound)
		}
		Msg = "用户输入的账号密码错误"
	}
	tpl, tplerr := template.ParseFiles([]string{"views/login/login.html"}...)
	if tplerr != nil {
		panic(tplerr)
	}
	tpl.ExecuteTemplate(w, "login.html", struct {
		name string
		Err  string
	}{name, Msg})

}

// 用户退出控制器
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	web.DeleteSession(w, r, web.LoadSession(w, r))
	http.Redirect(w, r, "/login/", http.StatusFound)

}

// 查询所有用户控制器
func QueryAllUserHandler(w http.ResponseWriter, r *http.Request) {
	data := services.QueryAllUsers()

	utils.ParseTemplate(w, r,
		[]string{"views/user/users.html"},
		"users.html",
		data,
	)
}

// 添加控制器
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == http.MethodPost {
		user := models.User{
			Name:   r.FormValue("name"),
			Sex:    utils.TextSex(r.FormValue("sex")),
			Passwd: utils.Md5sum(r.FormValue("passwd")),
			Addr:   r.FormValue("addr"),
			Tel:    r.FormValue("tel"),
		}
		services.AddUser(user)
		http.Redirect(w, r, "/user", http.StatusFound)
	}
	tpl, tplerr := template.ParseFiles([]string{"views/add/add.html"}...)
	if tplerr != nil {
		panic(tplerr)
	}
	tpl.ExecuteTemplate(w, "add.html", nil)
}

// 删除控制器
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	services.DeleteUser(r.FormValue("id"))
	http.Redirect(w, r, "/user", http.StatusFound)
}

// 变更控制器
func ModifyUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 如果是POST数据，则更新数据到DB
	if r.Method == http.MethodPost && len(r.PostForm) == 6 {
		user := models.User{
			ID:     utils.StrconvAtoi(r.FormValue("id")),
			Name:   r.FormValue("name"),
			Sex:    utils.TextSex(r.FormValue("sex")),
			Passwd: utils.Md5sum(r.FormValue("passwd")),
			Addr:   r.FormValue("addr"),
			Tel:    r.FormValue("tel"),
		}
		services.ModifyUser(user)
		http.Redirect(w, r, "/user", http.StatusFound)
	}

	// 根据获取的ID信息，检索数据库内部数据并显示到页面中
	id := r.FormValue("id")
	data, _ := services.QueryUserID(id)
	utils.ParseTemplate(w, r, []string{"views/modify/modify.html"},
		"modify.html",
		struct {
			ID     string
			Name   string
			Passwd string
			Sex    string
			Addr   string
			Tel    string
		}{utils.StrconvItoa(data.ID),
			data.Name,
			data.Passwd,
			utils.SexText(data.Sex),
			data.Addr,
			data.Tel,
		})
}
