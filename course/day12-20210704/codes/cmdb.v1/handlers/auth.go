package handlers

import (
	"cmdb/services"
	"cmdb/utils"
	"cmdb/web"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	token := web.LoadToken(w, r) // 加载token
	// 验证用户是否登陆
	if _, ok := token.Get("uid").(int64); ok {
		// 已登录, 跳转到登陆页面
		http.Redirect(w, r, "/users/", http.StatusFound)
		return
	}

	err := ""
	name := ""
	password := ""
	if r.Method == http.MethodPost {
		// 登陆验证
		name = r.PostFormValue("name")
		password = r.PostFormValue("password")

		if user := services.ValidateLogin(name, password); user != nil {
			// 登陆成功后跳转到 其他位置
			token.Set("uid", user.Id)  // 设置用户登陆状态
			web.DumpToken(w, r, token) // 持久化存储token到cookie
			http.Redirect(w, r, "/users/", http.StatusFound)
			return
		}
		err = "用户名或密码错误"
	}

	utils.ParseTemplate(w, r, []string{"views/auth/login.html"}, "login.html", struct {
		Name string
		Err  string
	}{name, err})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	web.DeleteToken(w, r)
	http.Redirect(w, r, "/login/", http.StatusFound)
}
