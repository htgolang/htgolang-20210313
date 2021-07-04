package controllers

import (
	"cmdb/services"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	AuthorizationController
}

func (c *UserController) Query() {
	web.ReadFromRequest(&c.Controller)

	c.TplName = "user/users.html"
	c.Data["Users"] = services.QueryUsers("")
}

func (c *UserController) Delete() {
	if id, err := c.GetInt64("id"); err == nil {
		services.DeleteUser(id)

		// flash消息存储
		flash := web.NewFlash()
		flash.Notice("删除成功")
		flash.Store(&c.Controller)

	}
	c.Redirect(web.URLFor("UserController.Query", "a", "b"), http.StatusFound)
}
