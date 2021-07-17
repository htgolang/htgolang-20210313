package controllers

import (
	"net/http"
	"usermanager/auth"

	"github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	web.Controller
}

func (this *LoginController) Get() {
	this.Data["Err"] = ""
	this.TplName = "auth/login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	user, err := auth.Authencation(username, password)
	if err != nil {
		this.Data["Err"] = err.Error()
		this.TplName = "auth/login.html"
		return
	}
	auth.Login(this.Ctx, user)
	this.Redirect("/user", http.StatusFound)
}

type LogoutController struct {
	AuthController
}

func (this *LogoutController) Get() {
	auth.Logout(this.Ctx)
	this.Redirect("/login", http.StatusFound)
}
