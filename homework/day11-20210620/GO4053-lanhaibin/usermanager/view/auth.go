package view

import (
	"net/http"
	"usermanager/auth"
	"usermanager/template"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		user, err := auth.Authencation(username, password)
		if err != nil {
			template.RenderTemplate(w, "login.html", struct {
				Name string
				Err  string
			}{Name: username, Err: err.Error()})
			return
		}

		auth.Login(w, r, user)
		http.Redirect(w, r, "/user", http.StatusFound)

	}
	template.RenderTemplate(w, "login.html", struct {
		Name string
		Err  string
	}{})
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	auth.Logout(w, r)
	http.Redirect(w, r, "/login", http.StatusFound)
}
