package routers

import (
	"cmdb/handlers"
	"net/http"
)

func init() {
	http.HandleFunc("/", handlers.QueryAllUserHandler)
	http.HandleFunc("/login/", handlers.LoginUserHandler)
	http.HandleFunc("/logout/", handlers.LogoutHandler)
	http.HandleFunc("/user/", handlers.QueryAllUserHandler)
	http.HandleFunc("/user/add/", handlers.AddUserHandler)
	http.HandleFunc("/user/delete/", handlers.DeleteUserHandler)
	http.HandleFunc("/user/modify/", handlers.ModifyUserHandler)

	http.HandleFunc("/api/v1/user/", handlers.APIQueryAllUserHandler)
	http.HandleFunc("/api/v1/user/query/", handlers.APIQueryUserHandler)
	http.HandleFunc("/api/v1/user/add/", handlers.APIAddUserHandler)
	http.HandleFunc("/api/v1/user/delete/", handlers.APIDeleteUserHandler)
	http.HandleFunc("/api/v1/user/modify/", handlers.APIModifyUserHandler)
}
