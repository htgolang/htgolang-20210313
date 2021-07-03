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
)

func QueryDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	currUid, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	currentUser := services.QueryUserByID(currUid)
	currentRole := services.QueryRoleByID(currentUser.RoleId)

	dps := services.QueryDepartment("")
	if dps != nil {
		err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/department/list.html"}, "list.html", struct {
			Departments []models.Department
			CurrentUser *models.User
			CurrentRole *models.Role
		}{dps, currentUser, currentRole})
		if err != nil {
			log.Println(err)
		}
	}
}

func DeleteDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	_, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	dpId, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	services.DeleteDepartment(dpId)
	http.Redirect(w, r, "/dp/list/", http.StatusFound)

}

func AddDepartmentHandler(w http.ResponseWriter, r *http.Request) {
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
		name := strings.TrimSpace(r.PostFormValue("dpname"))
		err := services.AddDepartment(name)
		if err != nil {
			errMsg = fmt.Sprintf("%s", err)
		} else {
			http.Redirect(w, r, "/dp/list/", http.StatusFound)
		}
	}

	err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/department/add.html"}, "add.html", struct {
		ErrMsg      string
		CurrentUser *models.User
		CurrentRole *models.Role
	}{errMsg, currentUser, currentRole})
	if err != nil {
		log.Println(err)
	}
}

func EditDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	currUid, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	currentUser := services.QueryUserByID(currUid)
	currentRole := services.QueryRoleByID(currentUser.RoleId)

	dp := &models.Department{}
	errMsg := ""
	if r.Method == http.MethodGet {
		dpid, _ := strconv.ParseInt(strings.TrimSpace(r.FormValue("id")), 10, 64)
		dp = services.QueryDepartmentByID(dpid)
	}

	if r.Method == http.MethodPost {
		dpid, _ := strconv.ParseInt(r.PostFormValue("dpid"), 10, 64)
		name := strings.TrimSpace(r.PostFormValue("dpname"))
		err := services.EditDepartment(dpid, name)
		if err != nil {
			errMsg = fmt.Sprintf("%s", err)
			dp.Id = dpid
			dp.Name = name
		} else {
			http.Redirect(w, r, "/dp/list/", http.StatusFound)
			return
		}
	}

	err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/department/edit.html"}, "edit.html", struct {
		Department  *models.Department
		ErrMsg      string
		CurrentUser *models.User
		CurrentRole *models.Role
	}{dp, errMsg, currentUser, currentRole})
	if err != nil {
		log.Println(err)
	}
}
