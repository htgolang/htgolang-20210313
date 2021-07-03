package routers

import (
	"cmdb/handlers"
	"net/http"
)

/*绑定 url和 handler*/

func init() {
	http.HandleFunc("/", handlers.PortalHandler)
	http.HandleFunc("/login/", handlers.LoginHandler)
	http.HandleFunc("/logout/", handlers.LogoutHandler)
	http.HandleFunc("/user/list/", handlers.QueryUsersHandler)
	http.HandleFunc("/user/delete/", handlers.DeleteUserHandler)
	http.HandleFunc("/user/add/", handlers.AddUserHandler)
	http.HandleFunc("/user/edit/", handlers.EditUserHandler)
	http.HandleFunc("/user/changepw/", handlers.ChangePwHandler)
	http.HandleFunc("/dp/list/", handlers.QueryDepartmentHandler)
	http.HandleFunc("/dp/add/", handlers.AddDepartmentHandler)
	http.HandleFunc("/dp/delete/", handlers.DeleteDepartmentHandler)
	http.HandleFunc("/dp/edit/", handlers.EditDepartmentHandler)
	http.HandleFunc("/asset/list/", handlers.QueryAssetHandler)
	http.HandleFunc("/asset/delete/", handlers.DeleteAssetHandler)
	http.HandleFunc("/asset/add/", handlers.AddAssettHandler)
	http.HandleFunc("/asset/edit/", handlers.EditAssetHandler)
	http.HandleFunc("/role/list/", handlers.ListRoleHandler)
}
