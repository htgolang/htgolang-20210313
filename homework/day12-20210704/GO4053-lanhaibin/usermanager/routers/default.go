package routers

import (
	"usermanager/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/login", new(controllers.LoginController))
	web.Router("/logout", new(controllers.LogoutController))

	web.Router("/user", new(controllers.UserQueryController))
	web.Router("/user/create", new(controllers.UserCreateController))
	web.Router("/user/modify", new(controllers.UserModifyController))
	web.Router("/user/delete", new(controllers.UserDeleteController))
	web.Router("/user/changepassword", new(controllers.UserPasswordController))

	web.Router("/department", new(controllers.DepartmentQueryController))
	web.Router("/department/create", new(controllers.DepartmentCreateController))
	web.Router("/department/delete", new(controllers.DepartmentDeleteController))
	web.Router("/department/modify", new(controllers.DepartmentModifyController))
}
