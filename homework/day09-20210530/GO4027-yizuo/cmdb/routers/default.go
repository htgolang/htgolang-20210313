package routers

import (
	"cmdb/handlers"
	"net/http"
)

func init() {
	http.HandleFunc("/user/", handlers.QueryAllUserHandler)
	http.HandleFunc("/user/query/", handlers.QueryUserHandler)
	http.HandleFunc("/user/add/", handlers.AddUserHandler)
	http.HandleFunc("/user/delete/", handlers.DeleteUserHandler)
	http.HandleFunc("/user/modify/", handlers.ModifyUserHandler)

}
