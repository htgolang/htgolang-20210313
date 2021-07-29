package routers

import (
	"cmdb/controllers"

	"github.com/beego/beego/v2/server/web"
)

/*绑定 url和 handler*/

func init() {
	// http.HandleFunc("/", handlers.PortalHandler)
	// http.HandleFunc("/login/", handlers.LoginHandler)
	// http.HandleFunc("/logout/", handlers.LogoutHandler)
	// http.HandleFunc("/user/list/", handlers.QueryUsersHandler)
	// http.HandleFunc("/user/delete/", handlers.DeleteUserHandler)
	// http.HandleFunc("/user/add/", handlers.AddUserHandler)
	// http.HandleFunc("/user/edit/", handlers.EditUserHandler)
	// http.HandleFunc("/user/changepw/", handlers.ChangePwHandler)
	// http.HandleFunc("/dp/list/", handlers.QueryDepartmentHandler)
	// http.HandleFunc("/dp/add/", handlers.AddDepartmentHandler)
	// http.HandleFunc("/dp/delete/", handlers.DeleteDepartmentHandler)
	// http.HandleFunc("/dp/edit/", handlers.EditDepartmentHandler)
	// http.HandleFunc("/asset/list/", handlers.QueryAssetHandler)
	// http.HandleFunc("/asset/delete/", handlers.DeleteAssetHandler)
	// http.HandleFunc("/asset/add/", handlers.AddAssettHandler)
	// http.HandleFunc("/asset/edit/", handlers.EditAssetHandler)
	// http.HandleFunc("/role/list/", handlers.ListRoleHandler)

	web.Router("/portal/", new(controllers.PortalController))
	web.AutoRouter(&controllers.AuthController{})
	web.AutoRouter(&controllers.UserController{})
	web.AutoRouter(&controllers.AssetController{})
	web.AutoRouter(&controllers.DepartmentController{})
	web.AutoRouter(&controllers.RoleController{})
}
