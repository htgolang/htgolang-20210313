package handlers

import (
	"cmdb/models"
	"cmdb/services"
	"cmdb/utils"
	"cmdb/web"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func QueryUsersHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	currUid, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	currentUser := services.QueryUserByID(currUid)
	currentRole := services.QueryRoleByID(currentUser.RoleId)

	keyword := r.FormValue("q")
	users := services.QueryUser(keyword)
	err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/user/list.html"}, "list.html", struct {
		Users       []models.User
		CurrentUser *models.User
		CurrentRole *models.Role
	}{users, currentUser, currentRole})
	if err != nil {
		log.Println(err)
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	_, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	services.DeleteUser(id)
	http.Redirect(w, r, "/user/list/", http.StatusFound)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	currUid, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	currentUser := services.QueryUserByID(currUid)
	currentRole := services.QueryRoleByID(currentUser.RoleId)

	var errMsg string
	roles := services.QueryRole()
	if r.Method == http.MethodPost {
		//获取提交信息
		username := strings.TrimSpace(r.PostFormValue("username"))
		sex := strings.TrimSpace(r.PostFormValue("sex"))
		birthday := strings.TrimSpace(r.PostFormValue("birthday"))
		tel := strings.TrimSpace(r.PostFormValue("tel"))
		email := strings.TrimSpace(r.PostFormValue("email"))
		addr := strings.TrimSpace(r.PostFormValue("addr"))
		description := strings.TrimSpace(r.PostFormValue("description"))
		password := r.PostFormValue("password")
		confirmpw := r.PostFormValue("confirmpw")
		roleId, _ := strconv.ParseInt(r.PostFormValue("roleid"), 10, 64)

		err := services.AddUser(username, sex, birthday, tel, email, addr, description, password, confirmpw, roleId)
		if err == nil {
			http.Redirect(w, r, "/user/list/", http.StatusFound)
			return
		} else {
			errMsg = fmt.Sprintf("%s", err)
		}
	}

	err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/user/add.html"}, "add.html", struct {
		Roles       []models.Role
		ErrMsg      string
		CurrentUser *models.User
		CurrentRole *models.Role
	}{roles, errMsg, currentUser, currentRole})
	if err != nil {
		log.Println(err)
	}
}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	currUid, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	currentUser := services.QueryUserByID(currUid)
	currentRole := services.QueryRoleByID(currentUser.RoleId)

	var errMsg string
	var user *models.User = &models.User{}
	roles := services.QueryRole()
	if r.Method == http.MethodGet {
		//获取用户id
		uid, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
		user = services.QueryUserByID(uid)
	}

	if r.Method == http.MethodPost {
		//获取提交信息
		uid, _ := strconv.ParseInt(r.PostFormValue("uid"), 10, 64)
		username := strings.TrimSpace(r.PostFormValue("username"))
		sex := strings.TrimSpace(r.PostFormValue("sex"))
		birthday := strings.TrimSpace(r.PostFormValue("birthday"))
		tel := strings.TrimSpace(r.PostFormValue("tel"))
		email := strings.TrimSpace(r.PostFormValue("email"))
		addr := strings.TrimSpace(r.PostFormValue("addr"))
		description := strings.TrimSpace(r.PostFormValue("description"))
		roleId, _ := strconv.ParseInt(r.PostFormValue("roleid"), 10, 64)

		err := services.EditUser(uid, username, sex, birthday, tel, email, addr, description, roleId)
		if err == nil {
			http.Redirect(w, r, "/user/list/", http.StatusFound)
			return
		} else {
			errMsg = fmt.Sprintf("%s", err)
			bir, _ := time.Parse("2006-01-02", birthday)
			user.Id = uid
			user.Name = username
			user.Sex = func() bool { return sex != "0" }()
			user.Birthday = &bir
			user.Telephone = tel
			user.Email = email
			user.Addr = addr
			user.Description = description
			user.RoleId = roleId
		}
	}

	if user != nil {
		err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/user/edit.html"}, "edit.html", struct {
			Roles       []models.Role
			User        *models.User
			ErrMsg      string
			CurrentUser *models.User
			CurrentRole *models.Role
		}{roles, user, errMsg, currentUser, currentRole})
		if err != nil {
			log.Println(err)
		}
	}
}

func ChangePwHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	currUid, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	currentUser := services.QueryUserByID(currUid)
	currentRole := services.QueryRoleByID(currentUser.RoleId)

	var errMsg string

	if r.Method == http.MethodPost {
		uid, _ := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("uid")), 10, 64)
		oldPassword := strings.TrimSpace(r.PostFormValue("oldpassword"))
		newPassword := strings.TrimSpace(r.PostFormValue("newpassword"))
		confirmpw := strings.TrimSpace(r.PostFormValue("confirmpw"))

		err := services.ChangePw(uid, oldPassword, newPassword, confirmpw)
		if err != nil {
			errMsg = fmt.Sprintf("%s", err)
		} else {
			//密码修改成功，重新登陆
			http.Redirect(w, r, "/logout/", http.StatusFound)
			return
		}
	}

	err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/user/changepw.html"}, "changepw.html", struct {
		ErrMsg      string
		CurrentUser *models.User
		CurrentRole *models.Role
	}{errMsg, currentUser, currentRole})
	if err != nil {
		log.Println(err)
		return
	}

}
