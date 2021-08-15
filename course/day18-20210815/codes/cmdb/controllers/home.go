package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	AuthorizationController
}

func (c *HomeController) Index() {
	c.Redirect(web.URLFor("UserController.Query"), http.StatusFound)
}
