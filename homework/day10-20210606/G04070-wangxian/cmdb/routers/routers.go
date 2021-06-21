package routers

import (
	"cmdb/handlers"
	"net/http"
)

/*绑定 url和 handler*/

func init() {
	http.HandleFunc("/login/", handlers.LoginHandler)
	http.HandleFunc("/logout/", handlers.LogoutHandler)
	http.HandleFunc("/user/list/", handlers.QueryUsersHandler)
	http.HandleFunc("/user/delete/", handlers.DeleteUserHandler)
	http.HandleFunc("/user/add/", handlers.AddUserHandler)
	http.HandleFunc("/user/edit/", handlers.EditUserHandler)
}
