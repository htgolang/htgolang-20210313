package handlers

import (
	"cmdb/models"
	"cmdb/services"
	"cmdb/utils"
	"net/http"
	"strconv"
)

func QueryUserHandler(uid int64, w http.ResponseWriter, r *http.Request) {
	users := services.QueryUsers("")
	currentUser := services.GetUserById(uid)
	// currentUser := &models.User{Name: "xxxx"}

	utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/user/users.html"}, "users.html", struct {
		CurrentUser *models.User
		Users       []models.User
	}{currentUser, users})
}

func DeleteUserHandler(uid int64, w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.ParseInt(r.FormValue("id"), 10, 64); err == nil {
		services.DeleteUser(id)
	}
	http.Redirect(w, r, "/users/", http.StatusFound)

}
