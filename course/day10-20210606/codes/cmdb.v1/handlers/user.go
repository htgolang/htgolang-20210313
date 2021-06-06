package handlers

import (
	"cmdb/services"
	"cmdb/utils"
	"cmdb/web"
	"net/http"
	"strconv"
)

func QueryUserHandler(w http.ResponseWriter, r *http.Request) {
	// 验证用户是否登陆
	session := web.LoadSession(w, r)
	if _, ok := session.Get("uid").(int); !ok {
		// 未登录, 跳转到登陆页面
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	users := services.QueryUsers("")
	utils.ParseTemplate(w, r, []string{"views/user/users.html"}, "users.html", users)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// 验证用户是否登陆
	session := web.LoadSession(w, r)
	if _, ok := session.Get("uid").(int); !ok {
		// 未登录, 跳转到登陆页面
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
		services.DeleteUser(id)
	}
	http.Redirect(w, r, "/users/", http.StatusFound)

}
