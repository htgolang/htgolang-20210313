package controllers

import "github.com/beego/beego/v2/server/web"

type UserController struct {
	web.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("User")
}
