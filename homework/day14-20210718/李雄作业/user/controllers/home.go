package controllers

import (
	base "user/base/controllers"
)

type HomeController struct {
	base.RequiredAuthController // 通过封装的 Controller控制器
}

func (c *HomeController) Index() {
	c.Ctx.WriteString("home")
}
