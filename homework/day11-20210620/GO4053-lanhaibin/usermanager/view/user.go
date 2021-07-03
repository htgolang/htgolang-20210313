package view

import (
	"net/http"
	"strconv"
	"usermanager/auth"
	"usermanager/model"
	"usermanager/service"
	"usermanager/template"
)

type response map[string]interface{}
type request map[string]string

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.PostForm.Get("name")
		brithday := r.PostForm.Get("brithday")
		sexs := r.PostForm.Get("sex")
		addr := r.PostForm.Get("addr")
		tel := r.PostForm.Get("tel")
		password := r.PostForm.Get("password")
		err := service.Us.Create(name, brithday, sexs == "1", addr, tel, password)
		if err != nil {
			template.RenderTemplate(w, "user_create.html", struct {
				Current model.User
				Err     string
			}{auth.GetCurrentUser(r), err.Error()})
			return
		}
		http.Redirect(w, r, "/user", http.StatusFound)
		return
	}
	template.RenderTemplate(w, "user_create.html", struct {
		Current model.User
		Err     string
	}{auth.GetCurrentUser(r), ""})

}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	ids := r.FormValue("id")
	id, _ := strconv.Atoi(ids)
	service.Us.Delete(id)

	http.Redirect(w, r, "/user", http.StatusFound)

}

func HandleQueryUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s := r.Form.Get("q")
	userList := service.Us.Query(s)
	template.RenderTemplate(w, "user_list.html", struct {
		Current model.User
		Users   []model.User
	}{auth.GetCurrentUser(r), userList})
}

func HandleModifyUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		r.ParseForm()
		ids := r.PostForm.Get("id")
		id, _ := strconv.Atoi(ids)
		name := r.PostForm.Get("name")
		brithday := r.PostForm.Get("brithday")
		sex := r.PostForm.Get("sex")
		addr := r.PostForm.Get("addr")
		tel := r.PostForm.Get("tel")

		service.Us.Modify(id, name, brithday, sex, addr, tel)
		http.Redirect(w, r, "/user", http.StatusFound)
	} else {
		ids := r.FormValue("id")
		id, _ := strconv.Atoi(ids)
		user, err := service.Us.Get(id)
		if err != nil {
			http.Redirect(w, r, "/user", http.StatusFound)
			return
		}
		template.RenderTemplate(w, "user_modify.html", struct {
			Current model.User
			User    model.User
		}{auth.GetCurrentUser(r), user})
	}
}

// func HandleUserList(w http.ResponseWriter, r *http.Request) {
// 	userList := service.Us.Query("")
// 	template.RenderTemplate(w, "user_list.html", userList)
// }

func HandleChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		oldpassword := r.PostForm.Get("oldpassword")
		newpassword := r.PostForm.Get("newpassword")
		confirmpassword := r.PostForm.Get("confirmpassword")
		user := auth.GetCurrentUser(r)
		if !user.CheckPassword(oldpassword) || newpassword != confirmpassword || len(newpassword) < 6 {
			template.RenderTemplate(w, "changepassword.html", struct {
				Current model.User
				Err     string
			}{user, "修改失败"})
			return
		}
		err := service.Us.ChangePassword(user.Id, newpassword)
		if err != nil {
			template.RenderTemplate(w, "changepassword.html", struct {
				Current model.User
				Err     string
			}{user, "修改失败"})
			return
		}
		auth.Logout(w, r)
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		template.RenderTemplate(w, "changepassword.html", struct {
			Current model.User
			Err     string
		}{auth.GetCurrentUser(r), ""})
	}
}
