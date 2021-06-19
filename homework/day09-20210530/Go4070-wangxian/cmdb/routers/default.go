package routers

import (
	"cmdb/handlers"
	"net/http"
)

func init() {
	http.HandleFunc("/user/query/", handlers.QueryUsersHandler)
	http.HandleFunc("/user/detail/", handlers.QueryUserDetailHandler)
	http.HandleFunc("/user/add/", handlers.AddUserHandler)
	http.HandleFunc("/user/delete/", handlers.DelUserHandler)
	http.HandleFunc("/user/edit/", handlers.EditUserHandler)
}
