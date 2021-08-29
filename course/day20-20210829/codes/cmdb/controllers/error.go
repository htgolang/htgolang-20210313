package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	web.Controller
}

func (c *ErrorController) ErrorDb() {
	c.TplName = "error/db.html"
}

func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
}
