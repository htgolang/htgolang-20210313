package controllers

import (
	"fmt"
	"time"
	"user/errors"
	"user/forms"
	"user/services"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

// 登录验证函数
func (c *AuthController) Login() {

	// 判断是否已经登录 存在了session 如果存在则直接跳转到主页,否则到登录页面
	if user := c.GetSession("user"); user != nil {
		c.Redirect("/user/query/", 302)
		return
	}
	// 打开页面,点击登录

	// var errMsg string
	form := &forms.LoginFrom{}

	// 错误信息对接 attr = []error 1个属性对应多个错误
	errors := errors.NewErrors()

	if c.Ctx.Input.IsPost() {
		// ParseForm 是解析传进来的参数 并序列化到forms.LoginFrom结构体中 此时LoginFrom结构体已经存在用户页面中提交的数据
		if err := c.ParseForm(form); err == nil {
			if user := services.Auth(form); user != nil {
				fmt.Println("登录成功 - ", time.Now())
				c.SetSession("user", user.ID)
				c.Redirect("/user/query", 302)
				errors.AddError("default", "登录成功")
				return
			} else {
				// 将错误信息添加到一个切片中，可以灵活的进行调用
				errors.AddError("default", "用户名或密码错误了")
				errors.AddError("username", "用户名错误")
			}
		}
	}
	// 清除错误信息 errors.Clear()
	c.Data["form"] = form
	c.Data["errors"] = errors
	c.TplName = "auth/login.html"
}

// 退出登录函数
func (c *AuthController) Logout() {
	fmt.Println("退出登录")
	// 销毁session
	c.DestroySession()
	// 重定向到登录页面
	c.Redirect("/auth/login", 302)
}
