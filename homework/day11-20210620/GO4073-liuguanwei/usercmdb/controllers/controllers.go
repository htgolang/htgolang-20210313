package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"usercmdb/models"
	"usercmdb/server"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/user.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	err = tpl.ExecuteTemplate(w, "user.html", server.GetUsers())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl, err := template.ParseFiles("template/create.html")
		if err != nil {
			w.WriteHeader(500)
			fmt.Println("/create", err)
			return
		}
		err = tpl.ExecuteTemplate(w, "create.html", nil)
		if err != nil {
			w.WriteHeader(500)
			fmt.Println("/create", err)
			return
		}
	}
	if r.Method == http.MethodPost {
		err := server.AddUser(
			r.FormValue("name"),
			r.FormValue("sex") == "true",
			r.FormValue("addr"),
			r.FormValue("tel"),
			r.FormValue("birthday"),
			r.FormValue("passwd"),
		)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		http.Redirect(w, r, "/", 301)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	err = server.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	http.Redirect(w, r, "/", 301)
}

func ModifyUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user, err := server.ParseUser(r)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		tpl, err := template.ParseFiles("template/modify.html")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}

		err = tpl.ExecuteTemplate(w, "modify.html",
			struct {
				*models.User
				Err string
			}{user, r.FormValue("err")})
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
	}

	if r.Method == http.MethodPost {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		user, err := server.ParseUser(r)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, fmt.Sprintf("/modify?id=%v&name=%v&sex=%v&tel=%v&birthday=%v&addr=%v&passwd=%v&err=%v", user.ID, user.Name, user.Sex, user.Tel, user.Birthday, user.Addr, user.Passwd, err.Error()), 302)
			return
		}
		fmt.Printf("Post method to get the user info: %#v\n", user)
		if err := server.ModifyAuth(user); err != nil {
			fmt.Println(err)
			http.Redirect(w, r, fmt.Sprintf("/modify?id=%v&name=%v&sex=%v&tel=%v&birthday=%v&addr=%v&passwd=%v&err=%v", user.ID, user.Name, user.Sex, user.Tel, user.Birthday, user.Addr, user.Passwd, err.Error()), 302)
			return
		}

		err = server.ModifyUser(user)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, fmt.Sprintf("/modify?id=%v&name=%v&sex=%v&tel=%v&birthday=%v&addr=%v&passwd=%v&err=%v", user.ID, user.Name, user.Sex, user.Tel, user.Birthday, user.Addr, user.Passwd, err.Error()), 302)
			return
		}
		http.Redirect(w, r, "/", 301)
	}
}
