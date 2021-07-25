package users

import (
	"github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	web.Controller
}

func (c *IndexController) Index() {
	c.Ctx.WriteString("users")
}
