package apis

import (
	"github.com/beego/beego/v2/server/web"
)

type ApiController struct {
	web.Controller
}

func (c *ApiController) Prepare() {
	// Header: X-Token: xxxx
	authToken := c.Ctx.Input.Header("x-token")
	if !c.isAuth(authToken) {
		c.Data["json"] = map[string]interface{}{
			"ok":     false,
			"reason": "token error",
		}
		c.ServeJSON()
	}
}

func (c *ApiController) isAuth(authToken string) bool {
	tokens := web.AppConfig.DefaultStrings("api::tokens", []string{})
	for _, token := range tokens {
		if authToken == token {
			return true
		}
	}
	return false
}
