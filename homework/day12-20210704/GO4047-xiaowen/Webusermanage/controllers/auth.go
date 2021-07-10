package controllers

import (
	"net/http"
	"webusermanage/forms"
	"webusermanage/models"
	"webusermanage/utils"

	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	web.Controller
}

func (c *AuthController) Login() {
	//判断是否已经登入
	if value := c.GetSession("uid"); value != nil {

		// c.Redirect("/user/query/", http.StatusFound)
		c.Redirect(web.URLFor("UserController.Query"), http.StatusFound)
		return
	}

	var form forms.Userform
	valid := validation.Validation{}
	// 验证登录
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(&form); err == nil {
			if b, err := valid.Valid(&form); err != nil {
				valid.SetError("flag", "用户提交数据错误")
			} else if b {
				c.SessionRegenerateID()
				c.SetSession("uid", form.User.Id)
				c.Redirect(web.URLFor("UserController.Query"), http.StatusFound)
				return
			}
		} else {
			valid.SetError("flag", "用户提交数据错误")
		}

	}
	c.TplName = "login/login.html"
	c.Data["errors"] = valid.ErrorMap()

}

func (c *AuthController) Logout() {
	c.SessionRegenerateID()
	c.DestroySession()
	c.Redirect(web.URLFor("AuthController.Login"), http.StatusFound)
}

type AuthorizationController struct {
	web.Controller
}

func (c *AuthorizationController) Prepare() {
	var currentUser *models.User
	if value := c.GetSession("uid"); value != nil {
		if uid, ok := value.(int); ok {
			currentUser = utils.GetUserById(uid)
		}
	}
	if currentUser == nil {
		c.Redirect(web.URLFor("AuthController.Login"), http.StatusFound)
		return
	}
	c.Data["CurrentUser"] = currentUser
}
