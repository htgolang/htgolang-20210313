package controllers

import (
	"cmdb/services"
	"net/http"

	"cmdb/forms"

	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	AuthorizationController
}

func (c *UserController) Query() {
	c.TplName = "user/users.html"
	c.Data["Users"] = services.UserService.Query(c.GetString("q"))
}

func (c *UserController) Create() {
	// Get请求 打开页面
	valid := &validation.Validation{}
	form := &forms.CreateUserForm{}
	if c.Ctx.Input.IsPost() {
		// Post请求 创建用户
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("default", "提交数据错误")
			} else if ok {
				// 存储数据
				services.UserService.Create(form.User) // 根据存储结果给用户响应

				c.Redirect(web.URLFor("UserController.Query"), http.StatusFound)
			}
		} else {
			valid.SetError("default", "提交数据错误")
		}
	}

	c.Data["form"] = form
	c.Data["errors"] = valid.ErrorMap()
	c.TplName = "user/create.html"
}

func (c *UserController) ModifyPassword() {
	var success = ""
	valid := &validation.Validation{}
	form := &forms.ModifyPasswordForm{CurrentUser: c.CurrentUser}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("default", "提交数据错误")
			} else if ok {
				// 更新密码
				c.CurrentUser.SetPassword(form.NewPassword)
				services.UserService.ModifyPassword(c.CurrentUser)
				success = "修改密码成功, 请重新登陆"
				c.DestroySession() // 清空session 重新登陆
			}
		} else {
			valid.SetError("default", "提交数据错误")
		}
	}
	c.Data["form"] = form
	c.Data["errors"] = valid.ErrorMap()
	c.Data["success"] = success

	c.TplName = "user/modify_password.html"
}

func (c *UserController) Delete() {
	if id, err := c.GetInt64("id"); err == nil {
		services.UserService.Delete(id)
	}
	c.Redirect(web.URLFor("UserController.Query"), http.StatusFound)
}
