package controllers

import (
	"usermanager/auth"
	"usermanager/model"

	"github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	web.Controller
	User model.User
}

func (this *AuthController) Prepare() {
	auth.LoginRequired(this.Ctx)
	this.User = auth.GetCurrentUser(this.Ctx)
	this.Data["Current"] = this.User
}
