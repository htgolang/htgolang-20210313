package view

import (
	"fmt"
	"net/http"
	"strconv"
	"usermanager/service"
	"usermanager/template"
)

type response map[string]interface{}
type request map[string]string

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.PostForm.Get("name")
		addr := r.PostForm.Get("addr")
		tel := r.PostForm.Get("tel")
		password := r.PostForm.Get("password")
		err := service.Us.Create(name, addr, tel, password)
		if err != nil {
			template.RenderTemplate(w, "user_create.html", err.Error())
			return
		}
		http.Redirect(w, r, "/user", http.StatusFound)
		return
	}
	template.RenderTemplate(w, "user_create.html", "")

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
	fmt.Println(s)
	userList := service.Us.Query(s)
	template.RenderTemplate(w, "user_list.html", userList)
}

func HandleModifyUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		r.ParseForm()
		ids := r.PostForm.Get("id")
		id, _ := strconv.Atoi(ids)
		name := r.PostForm.Get("name")
		addr := r.PostForm.Get("addr")
		tel := r.PostForm.Get("tel")

		service.Us.Modify(id, name, addr, tel)
		http.Redirect(w, r, "/user", http.StatusFound)
	} else {
		ids := r.FormValue("id")
		id, _ := strconv.Atoi(ids)
		user, err := service.Us.Get(id)
		if err != nil {
			http.Redirect(w, r, "/user", http.StatusFound)
			return
		}
		template.RenderTemplate(w, "user_modify.html", user)
	}
}

func HandleUserList(w http.ResponseWriter, r *http.Request) {
	userList := service.Us.Query("")
	fmt.Println(userList)
	template.RenderTemplate(w, "user_list.html", userList)
}
