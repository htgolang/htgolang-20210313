package controllers

import (
	"cmdb/forms"
	"cmdb/models"
	"cmdb/services"
	"net/http"

	"github.com/beego/beego/v2/adapter/validation"
	"github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	web.Controller
}

func (c *AuthController) Login() {
	//验证是否登录，已登录跳转到user/list
	if value := c.GetSession("uid"); value != nil {
		if _, ok := value.(int64); ok {
			c.Redirect(web.URLFor("UserController.Query"), http.StatusFound)
			return
		}
	}

	var form forms.LoginForm
	var valid validation.Validation
	if c.Ctx.Input.IsPost() {
		//获取用户名密码
		if err := c.ParseForm(&form); err == nil {
			if b, err := valid.Valid(&form); err != nil {
				valid.SetError("user", "提交数据错误")
			} else if b {
				//验证通过，设置session, 保存session
				c.SetSession("uid", form.User.Id)

				//跳转到portal页面
				c.Redirect("/portal/", http.StatusFound)
				return
			}
		} else {
			valid.SetError("user", "提交数据错误")
		}
	}
	c.Data["ErrMsgs"] = valid.ErrorMap()
	c.TplName = "auth/login.html"

}

func (c *AuthController) Logout() {
	//清除session
	c.DestroySession()

	//跳转登录界面
	c.Redirect(web.URLFor("AuthController.Login"), http.StatusFound)
}

type AuthenticationController struct {
	web.Controller
	CurrentUser *models.User
}

func (c *AuthenticationController) Prepare() {
	//验证是否登录
	var currentUser *models.User

	if value := c.GetSession("uid"); value != nil {
		if uid, ok := value.(int64); ok {
			currentUser = services.QueryUserByID(uid)
		}
	}

	if currentUser == nil {
		c.Redirect(web.URLFor("AuthController.Login"), http.StatusFound)
		return
	}

	currentRole := services.QueryRoleByID(currentUser.RoleId)
	c.Data["CurrentUser"] = currentUser
	c.Data["CurrentRole"] = currentRole
	c.CurrentUser = currentUser
}
