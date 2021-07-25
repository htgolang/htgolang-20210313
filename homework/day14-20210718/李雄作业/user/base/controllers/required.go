package controllers

import (
	"fmt"
	"user/services"

	"github.com/astaxie/beego"
)

type RequiredAuthController struct {
	beego.Controller
}

// 控制器添加功能
func (c *RequiredAuthController) Prepare() {
	c.Data["navKey"] = ""
	user := c.GetSession("user")
	if user == nil {
		// 为空则未找到session数据库，状态为未登录
		fmt.Println("用户未登录,重定向到登录页面")
		c.Redirect("/auth/login", 302)
		return
	}
	// 空接口，断言
	if pk, ok := user.(int64); ok {
		if user := services.GetUserById(pk); user != nil {
			c.Data["currentUser"] = user
		}
	}

	if c.Data["currentUser"] == nil {
		// 清空session 方式死循环跳转
		c.DestroySession()
		c.Redirect("/auth/login", 302)
		return
	}
}
