package handlers

import (
	"cmdb/models"
	"cmdb/services"
	"cmdb/utils"
	"cmdb/web"
	"log"
	"net/http"
)

func PortalHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	currUid, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	currentUser := services.QueryUserByID(currUid)
	currentRole := services.QueryRoleByID(currentUser.RoleId)

	err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/base/portal.html"}, "portal.html", struct {
		CurrentUser *models.User
		CurrentRole *models.Role
	}{currentUser, currentRole})
	if err != nil {
		log.Println(err)
		return
	}

}
