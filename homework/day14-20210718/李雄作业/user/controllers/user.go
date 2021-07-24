package controllers

import (
	"fmt"
	base "user/base/controllers"
	"user/forms"
	"user/services"
)

type UserController struct {
	base.RequiredAuthController // 封装了 Prepare函数的controller
}

// 会在所有的函数执行前进行检查
// func (c *UserController) Prepare() {
// 	user := c.GetSession("user")
// 	if user == nil {
// 		// session等于nil 那就属于未登录,然后重定向到登录页面
// 		c.Redirect("/auth/login", 302)
// 		return // 保证后面的代码不执行
// 	}
// }

// 菜单选中
func (c *UserController) Prepare() {
	c.RequiredAuthController.Prepare()
	c.Data["navKey"] = "user"
}

// 添加函数
func (c *UserController) Add() {
	if c.Ctx.Input.IsPost() {
		var form forms.AddUserForm
		c.ParseForm(&form)
		services.AddUser(form.Name, form.Password, form.Addr, form.Sex)
		c.Redirect("/user/query", 302)
	} else {
		c.Data["form"] = &forms.AddUserForm{Name: ""}
		// c.Data["navKey"] = "user"
		c.TplName = "user/add.html"
	}
}

// 查询函数
func (c *UserController) Query() {
	// c.Data["navKey"] = "user"
	c.Data["users"] = services.GetUsers()
	c.TplName = "user/query.html"
}

// 删除函数
func (c *UserController) Delete() {
	if id, err := c.GetInt64("id"); err == nil {
		services.DeleteUser(id)
	}
	c.Redirect("/user/query", 302)
}

// 修改密码
func (c *UserController) Update() {
	if c.Ctx.Input.IsPost() {
		var form forms.ChangeUserForm
		c.ParseForm(&form)
		fmt.Println(form.Name, form.Password, form.Addr, form.Sex)
		services.ChangeUser(form.Name, form.Password, form.Addr, form.Sex)
		c.Redirect("/user/query/", 302)
	} else {
		c.TplName = "user/update.html"
	}
}
