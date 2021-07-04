package main

import (
	"github.com/beego/beego/v2/server/web/context"

	"github.com/beego/beego/v2/server/web"
)

type RestfulController struct {
	web.Controller
}

func (c *RestfulController) Get() {
	c.Ctx.WriteString("get:" + c.Ctx.Input.Param(":id"))
}

type CustomController struct {
	web.Controller
}

func (c *CustomController) Open() {
	c.Ctx.WriteString("open:" + c.Ctx.Request.Method)
}

func (c *CustomController) Update() {
	c.Ctx.WriteString("update:" + c.Ctx.Request.Method)
}

func (c *CustomController) Close() {
	c.Ctx.WriteString("close:" + c.Ctx.Request.Method)
}

type AuthController struct {
	web.Controller
}

// /auth/login/
func (c *AuthController) Login() {
	c.Ctx.WriteString("login:" + c.Ctx.Request.Method)
}

// /auth/logout/
func (c *AuthController) Logout() {
	c.Ctx.WriteString("logout:" + c.Ctx.Request.Method)
}

func main() {
	// url -> handler/handlerfunc
	// url -> Controller/ControllerFunc

	// 定义路由函数

	// 指定处理器函数(控制器函数) Ctx
	// web.Get/Post/Put/Delete/Head/Options/Patch/Any
	// web.Get
	web.Get("/", func(ctx *context.Context) {
		ctx.WriteString("get")
	})
	// web.Post
	// web.Post("/", func(ctx *context.Context) {
	// 	ctx.WriteString("post")
	// })
	// web.Delete
	web.Delete("/", func(ctx *context.Context) {
		ctx.WriteString("delete")
	})
	// web.Head
	web.Head("/", func(ctx *context.Context) {
		ctx.WriteString("head")
	})
	// web.Put
	web.Put("/", func(ctx *context.Context) {
		ctx.WriteString("put")
	})

	web.Any("/any/", func(ctx *context.Context) {
		ctx.WriteString("any:" + ctx.Request.Method)
	})

	// 正则路由
	// /regex/
	// /regex/1/
	// /regex/2/
	// /regex/3/
	// 正则 数字: [0-9]
	web.Get("/regex/?:id([0-9]+)/", func(ctx *context.Context) {
		ctx.WriteString(ctx.Input.Param(":id"))
	})

	web.Get("/regexint/?:id:int/", func(ctx *context.Context) {
		ctx.WriteString(ctx.Input.Param(":id"))
	})

	web.Get("/regexpath/*", func(ctx *context.Context) {
		ctx.WriteString(ctx.Input.Param(":splat"))
	})

	web.Get("/regexext/*.*", func(ctx *context.Context) {
		ctx.WriteString(ctx.Input.Param(":path") + ":" + ctx.Input.Param(":ext"))
	})
	// 指定处理器(控制器)

	// resultful/id/
	web.Router("/restful/", new(RestfulController))
	web.Router("/restful/:id:int/", new(RestfulController))

	// Get/Post Get=>Open Post=> Update Delete=>Update  => Close
	// requestmethod:funcmethod;
	web.Router("/custom/", new(CustomController), "get:Open;post,delete:Update;*:Close")

	// 自动路由
	// url => controllerName:methodName

	// web.Router("/auth/login/", new(AuthController), "*:Login")
	// web.Router("/auth/logout/", new(AuthController), "*:Logout")

	// AuthController Login Logout方法
	// /auth/login => Login
	// /auth/logout => Logout
	web.AutoRouter(new(AuthController))

	// 注解路由
	// 不要用 => 只能在开发模式使用 => 生成go文件

	web.Run()
}
