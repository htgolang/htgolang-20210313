package apis

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type ApiController struct {
	web.Controller
}

func (c *ApiController) authToken() string {
	// Agent
	token := c.Ctx.Input.Header("x-token")
	if token != "" {
		return token
	}

	// alertmanager
	token = c.Ctx.Input.Header("Authorization")
	if token != "" {
		elements := strings.SplitN(token, " ", 2)
		return elements[len(elements)-1]
	}

	return ""
}

func (c *ApiController) Prepare() {
	if !c.isAuth(c.authToken()) {
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
