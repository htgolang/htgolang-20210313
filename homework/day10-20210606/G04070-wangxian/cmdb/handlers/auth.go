package handlers

import (
	"cmdb/services"
	"cmdb/utils"
	"cmdb/web"
	"log"
	"net/http"
)

/*定义handler*/

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	errMsg := ""

	//验证用户是否登录
	session := web.LoadSession(w, r)
	if _, ok := session.Get("uid").(int); ok {
		//已经登陆,跳转到/users/
		http.Redirect(w, r, "/user/list/", http.StatusFound)
	}

	if r.Method == http.MethodPost {
		//获取用户名密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("passwd")

		//验证用户名密码
		if user := services.Login(username, password); user != nil {
			//验证通过，设置session, 保存session
			session.Set("uid", 1)
			web.DumpSession(w, r, session)
			//跳转到user-list页面
			http.Redirect(w, r, "/user/list/", http.StatusFound)
			return
		}
		errMsg = "用户名密码错误"
	}

	err := utils.ParseTemplate(w, r, []string{"views/auth/login.html"}, "login.html", errMsg)
	if err != nil {
		log.Println(err)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	web.DeleteSession(w, r, web.LoadSession(w, r))

	http.Redirect(w, r, "/login/", http.StatusFound)
}
