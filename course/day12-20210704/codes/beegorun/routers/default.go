package routers

import (
	"beegorun/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// 定义路由
	web.Router("/", &controllers.IndexController{}) // http.Handle
	web.Router("/user/", &controllers.UserController{})
}
