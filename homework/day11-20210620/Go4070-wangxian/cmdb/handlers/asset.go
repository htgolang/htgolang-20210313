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

func QueryAssetHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	currUid, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	currentUser := services.QueryUserByID(currUid)
	currentRole := services.QueryRoleByID(currentUser.RoleId)

	ats := services.QueryAsset()
	err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/asset/list.html"}, "list.html", struct {
		Assets      []models.Asset
		CurrentUser *models.User
		CurrentRole *models.Role
	}{ats, currentUser, currentRole})
	if err != nil {
		log.Println(err)
	}

}

func DeleteAssetHandler(w http.ResponseWriter, r *http.Request) {
	session := web.LoadSession(w, r)
	_, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	aid, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	services.DleteAsset(aid)
	http.Redirect(w, r, "/asset/list/", http.StatusFound)
}

func AddAssettHandler(w http.ResponseWriter, r *http.Request) {
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
		ip := strings.TrimSpace(r.PostFormValue("ip"))
		addr := strings.TrimSpace(r.PostFormValue("addr"))
		err := services.AddAsset(ip, addr)
		if err != nil {
			errMsg = fmt.Sprintf("%s", err)
		} else {
			http.Redirect(w, r, "/asset/list/", http.StatusFound)
		}
	}

	err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/asset/add.html"}, "add.html", struct {
		ErrMsg      string
		CurrentUser *models.User
		CurrentRole *models.Role
	}{errMsg, currentUser, currentRole})
	if err != nil {
		log.Println(err)
	}
}

func EditAssetHandler(w http.ResponseWriter, r *http.Request) {
	session := web.LoadSession(w, r)
	currUid, ok := session.Get("uid").(int64)
	if !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}
	currentUser := services.QueryUserByID(currUid)
	currentRole := services.QueryRoleByID(currentUser.RoleId)

	asset := &models.Asset{}
	errMsg := ""
	if r.Method == http.MethodGet {
		aid, _ := strconv.ParseInt(strings.TrimSpace(r.FormValue("id")), 10, 64)
		asset = services.QueryAssetByID(aid)
	}

	if r.Method == http.MethodPost {
		aid, _ := strconv.ParseInt(r.PostFormValue("aid"), 10, 64)
		ip := strings.TrimSpace(r.PostFormValue("ip"))
		addr := strings.TrimSpace(r.PostFormValue("addr"))
		err := services.EditAsset(aid, ip, addr)
		if err != nil {
			errMsg = fmt.Sprintf("%s", err)
			asset.Id = aid
			asset.Ip = ip
			asset.Addr = addr
		} else {
			http.Redirect(w, r, "/asset/list/", http.StatusFound)
			return
		}
	}

	err := utils.ParseTemplate(w, r, []string{"views/base/header.html", "views/asset/edit.html"}, "edit.html", struct {
		Asset       *models.Asset
		ErrMsg      string
		CurrentUser *models.User
		CurrentRole *models.Role
	}{asset, errMsg, currentUser, currentRole})
	if err != nil {
		log.Println(err)
	}
}
