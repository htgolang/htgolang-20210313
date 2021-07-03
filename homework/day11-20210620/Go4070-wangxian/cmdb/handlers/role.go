package handlers

import (
	"cmdb/models"
	"cmdb/services"
	"cmdb/utils"
	"cmdb/web"
	"log"
	"net/http"
)

func ListRoleHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	currUid, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	currentUser := services.QueryUserByID(currUid)
	currentRole := services.QueryRoleByID(currentUser.RoleId)

	roles := services.QueryRole()
	err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/role/list.html"}, "list.html", struct {
		Roles       []models.Role
		CurrentUser *models.User
		CurrentRole *models.Role
	}{roles, currentUser, currentRole})
	if err != nil {
		log.Println(err)
	}
}
