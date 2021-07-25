package routers

import (
	"user/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 用户登录
	beego.AutoRouter(&controllers.AuthController{}) // 自动注册路由 auth/login
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.HomeController{}) // 主页路由
}
