package routers

import (
	"cmdb/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.ErrorController(new(controllers.ErrorController))

	web.Router("/", new(controllers.HomeController), "*:Index")
	web.AutoRouter(new(controllers.AuthController))
	web.AutoRouter(new(controllers.UserController))
	web.AutoRouter(new(controllers.AssetController))
}
