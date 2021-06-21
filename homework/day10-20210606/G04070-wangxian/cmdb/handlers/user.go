package handlers

import (
	"cmdb/modules"
	"cmdb/services"
	"cmdb/utils"
	"cmdb/web"
	"log"
	"net/http"
	"strconv"
)

func QueryUsersHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	// fmt.Println(session)
	if _, ok := session.Get("uid").(int); !ok {
		// fmt.Println(ok)
		// fmt.Printf("%T", session.Get("uid"))
		//未登陆,跳转到/login/
		http.Redirect(w, r, "/login/", http.StatusFound)
	}

	users := services.QueryUser("")
	err := utils.ParseTemplate(w, r, []string{"views/user/list.html"}, "list.html", users)
	if err != nil {
		log.Println(err)
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	if _, ok := session.Get("uid").(int); !ok {
		//未登陆,跳转到/login/
		http.Redirect(w, r, "/login/", http.StatusFound)
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	services.DeleteUser(id)
	http.Redirect(w, r, "/user/list/", http.StatusFound)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	if _, ok := session.Get("uid").(int); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
	}

	errMsg := ""
	if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		sex := r.PostFormValue("sex")
		addr := r.PostFormValue("addr")
		tel := r.PostFormValue("tel")
		password := r.PostFormValue("password")
		confirmpw := r.PostFormValue("confirmpw")
		errMsg = services.AddUser(username, sex, addr, tel, password, confirmpw)
		if errMsg == "" {
			http.Redirect(w, r, "/user/list/", http.StatusFound)
			return
		}
	}

	err := utils.ParseTemplate(w, r, []string{"views/user/add.html"}, "add.html", errMsg)
	if err != nil {
		log.Println(err)
	}
}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	//验证用户是否登录
	session := web.LoadSession(w, r)
	if _, ok := session.Get("uid").(int); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
	}

	if r.Method == http.MethodPost {
		uid, _ := strconv.Atoi(r.PostFormValue("uid"))
		username := r.PostFormValue("username")
		sex := r.PostFormValue("sex")
		addr := r.PostFormValue("addr")
		tel := r.PostFormValue("tel")

		errMsg := services.EditUser(uid, username, sex, addr, tel)
		if errMsg == "" {
			http.Redirect(w, r, "/user/list/", http.StatusFound)
			return
		} else {
			err := utils.ParseTemplate(w, r, []string{"views/user/edit.html"}, "edit.html", struct {
				Data   *modules.User
				ErrMsg string
			}{&modules.User{Id: uid, Name: username, Sex: func() bool { return sex != "0" }(), Tel: tel, Addr: addr}, errMsg})
			if err != nil {
				log.Println(err)
			}
			return
		}

	}

	//获取用户id
	uid, _ := strconv.Atoi(r.FormValue("id"))
	user := services.QueryUserByID(uid)
	if user != nil {
		err := utils.ParseTemplate(w, r, []string{"views/user/edit.html"}, "edit.html", struct {
			Data   *modules.User
			ErrMsg string
		}{user, ""})
		if err != nil {
			log.Println(err)
		}
	}
}
