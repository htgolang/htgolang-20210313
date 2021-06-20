package handlers

import (
	"cmdb/models"
	"cmdb/services"
	"cmdb/utils"
	"cmdb/web"
	"net/http"
	"strconv"
)

func QueryUserHandler(w http.ResponseWriter, r *http.Request) {
	// 验证用户是否登陆
	token := web.LoadToken(w, r) // 加载token
	uid, ok := token.Get("uid").(int64)

	if !ok {
		// 已登录, 跳转到登陆页面
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	users := services.QueryUsers("")
	currentUser := services.GetUserById(uid)

	utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/user/users.html"}, "users.html", struct {
		CurrentUser *models.User
		Users       []models.User
	}{currentUser, users})
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// 验证用户是否登陆
	token := web.LoadToken(w, r) // 加载token
	if _, ok := token.Get("uid").(int64); !ok {
		// 已登录, 跳转到登陆页面
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	if id, err := strconv.ParseInt(r.FormValue("id"), 10, 64); err == nil {
		services.DeleteUser(id)
	}
	http.Redirect(w, r, "/users/", http.StatusFound)

}
