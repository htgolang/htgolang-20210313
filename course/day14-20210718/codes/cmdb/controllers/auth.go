package controllers

import (
	"net/http"

	"cmdb/forms"
	"cmdb/models"
	"cmdb/services"

	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	web.Controller
}

func (c *AuthController) Login() {
	// 若已经登陆则跳转到用户列表页面
	if value := c.GetSession("uid"); value != nil {
		if _, ok := value.(int64); ok {
			// 可以再检查一下 id是否再数据库中存储(状态是否正常)
			c.Redirect(web.URLFor("HomeController.Index"), http.StatusFound)
			return
		}
	}

	var form forms.LoginForm
	valid := validation.Validation{}

	if c.Ctx.Input.IsPost() {
		// 登陆验证
		if err := c.ParseForm(&form); err == nil {
			if b, err := valid.Valid(&form); err != nil {
				valid.SetError("user", "提交数据错误")
			} else if b {
				c.SessionRegenerateID()
				c.SetSession("uid", form.User.Id)
				c.Redirect(web.URLFor("HomeController.Index"), http.StatusFound)
				return
			}
		} else {
			valid.SetError("user", "提交数据错误")
		}
	}
	c.TplName = "auth/login.html"
	c.Data["form"] = form
	c.Data["errors"] = valid.ErrorMap()
}

func (c *AuthController) Logout() {
	// 销毁session
	c.DestroySession()
	c.SessionRegenerateID()
	c.Redirect(web.URLFor("AuthController.Login"), http.StatusFound)
}

type AuthorizationController struct {
	web.Controller

	CurrentUser *models.User
}

func (c *AuthorizationController) Prepare() {
	var currentUser *models.User

	if value := c.GetSession("uid"); value != nil {
		if uid, ok := value.(int64); ok {
			currentUser = services.UserService.GetById(uid)
		}
	}

	if currentUser == nil {
		c.Redirect(web.URLFor("AuthController.Login"), http.StatusFound)
		return
	}
	// 校验权限
	controllerName, actionName := c.GetControllerAndAction()

	if !services.RoleService.IsPermission(currentUser.RoleId, controllerName, actionName) {
		c.Abort("Permission")
		return
	}

	// 查询 用户所属角色具有的模块信息
	c.Data["CurrentUser"] = currentUser
	c.Data["Menus"] = services.RoleService.GetMenus(currentUser.RoleId)
	c.CurrentUser = currentUser
}
