package main

import (
	"net/http"
	"strconv"
	"text/template"
)

type User struct {
	Id   int
	Name string
	Sex  bool
}

//get the user id
func GetId(users []*User) int {
	id := 0
	for _, user := range users {
		if user.Id > id {
			id = user.Id
		}
	}
	return id + 1
}

//get the user
func GetUser(users []*User, id string) *User {
	if uid, err := strconv.Atoi(id); err == nil {
		for _, user := range users {
			if user.Id == uid {
				return user
			}
		}
	}
	return nil
}

func main() {
	addr := ":9999"
	//base users
	users := []*User{
		{1, "reid", true},
		{2, "mark", true},
		{3, "mary", false},
		{4, "mike", true},
	}

	//users list
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("templates/users.html"))
		tpl.ExecuteTemplate(w, "users.html", users)
	})

	//delete user
	http.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
		if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
			tempUsers := make([]*User, 0, len(users))
			for _, user := range users {
				if user.Id != id {
					tempUsers = append(tempUsers, user)
				}
				users = tempUsers
			}
		}
		http.Redirect(w, r, "/", 302)
	})

	//add user
	http.HandleFunc("/add/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tpl := template.Must(template.ParseFiles("templates/add.html"))
			tpl.ExecuteTemplate(w, "add.html", nil)
		} else {
			users = append(users, &User{
				GetId(users),
				r.FormValue("name"),
				r.FormValue("sex") == "1",
			})
			http.Redirect(w, r, "/", 302)
		}
	})

	//modify user
	http.HandleFunc("/modify/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			user := GetUser(users, r.FormValue("id"))
			tpl := template.Must(template.ParseFiles("templates/modify.html"))
			tpl.ExecuteTemplate(w, "modify.html", user)
		} else {
			if uid, err := strconv.Atoi(r.FormValue("id")); err == nil {
				for k, u := range users {
					if u.Id == uid {
						users[k] = &User{
							uid,
							r.FormValue("name"),
							r.FormValue("sex") == "1",
						}
					}
				}
			}
			http.Redirect(w, r, "/", 302)
		}
	})
	http.ListenAndServe(addr, nil)
}
