package main

import (
	"namespace/controllers/roles"
	"namespace/controllers/users"

	"github.com/beego/beego/v2/server/web/context"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	// web.AutoRouter(new(users.IndexController))
	// web.AutoRouter(new(roles.IndexController))
	// web.Get/web.Post/...
	// web.Router

	// web.NSGet/NSPort
	// web.NSRouter

	userNamespace := web.NewNamespace("/users",
		// /users/index/index
		web.NSAutoRouter(new(users.IndexController)),

		// /users/get
		web.NSGet("/get", func(c *context.Context) {
			c.WriteString("users.get")
		}),
	)

	roleNamespace := web.NewNamespace("/roles",
		// /users/index/index
		web.NSAutoRouter(new(roles.IndexController)),
	)

	web.AddNamespace(userNamespace)
	web.AddNamespace(roleNamespace)

	// controller.action => url
	// 在Controller.Action冲突场景 不能使用urlfor/web.URLFor 写URL
	// web.URLFor // urlfor
	// web.URLFor("IndexController.Index")
	// web.URLFor("IndexController.Index")

	// v1
	// Query()
	// v2
	// Query()
	web.Run()
}
