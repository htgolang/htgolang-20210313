package main

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type RequestController struct {
	web.Controller
}

func (c *RequestController) Index() {
	// c.Ctx.Request => http.Request
	// c.Ctx.Input => Query/Bind header
	// c.Ctx => Cookie
	// Controller => GetXXX Input ParseForm SaveToFile
	// 从上到下 为封装层次
	// 使用 优先是由 封装层次最多的 （封装层次越多，提供的特性越多）

	// 登陆 username password

	fmt.Println("username:", c.GetString("username"))
	fmt.Println("password:", c.GetString("password"))

	c.Ctx.WriteString("index")
}

func (c *RequestController) Bind() {
	// Scan
	var name string
	var password string
	c.Ctx.Input.Bind(&name, "username")
	c.Ctx.Input.Bind(&password, "password")

	fmt.Println(name, ":", password)
	c.Ctx.WriteString("bind")
}

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// 提交数据 解析到对象(类似 数据库行记录 对应到model的对象)
func (c *RequestController) Parse() {
	var form LoginForm
	err := c.ParseForm(&form)
	fmt.Printf("%v, %#v\n", err, form)

	c.Ctx.WriteString("parse")
}

func (c *RequestController) Json() {
	body := c.Ctx.Input.CopyBody(1024 * 1024)
	// 自动配置 app.conf // copyrequestbody=true
	fmt.Println(string(body))
	fmt.Println(string(c.Ctx.Input.RequestBody))
	c.Ctx.WriteString("json")
}

func main() {
	// 控制器函数 ctx *context.Context
	// 控制器 Ctx *context.Context
	web.AutoRouter(new(RequestController))
	web.Run()
}
