package routers

import (
	"cmdb/handlers"
	"net/http"
)

func init() {
	http.HandleFunc("/login/", handlers.LoginHandler)
	http.HandleFunc("/logout/", handlers.LogoutHandler)
	http.HandleFunc("/users/", handlers.QueryUserHandler)
	http.HandleFunc("/user/delete/", handlers.DeleteUserHandler)
}
