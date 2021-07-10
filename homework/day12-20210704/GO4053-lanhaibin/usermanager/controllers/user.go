package controllers

import (
	"fmt"
	"net/http"
	"usermanager/auth"
	"usermanager/forms"
	"usermanager/service"

	"github.com/beego/beego/v2/core/validation"
)

type UserCreateController struct {
	AuthController
}

func (this *UserCreateController) Get() {
	this.Data["Err"] = ""
	this.TplName = "user/user_create.html"
}

func (this *UserCreateController) Post() {
	// name := this.GetString("name")
	// brithday := this.GetString("brithday")
	// sexs := this.GetString("sex")
	// addr := this.GetString("addr")
	// tel := this.GetString("tel")
	// password := this.GetString("password")
	valid := validation.Validation{}
	form := forms.UserCreateForm{}
	this.ParseForm(&form)
	_, err := valid.Valid(&form)
	if err != nil {
		fmt.Println(err)
		this.Data["Err"] = err.Error()
		this.TplName = "user/user_create.html"
		return
	}
	err = service.Us.Create(
		form.Name,
		form.Brithday,
		form.Sex == "1",
		form.Addr,
		form.Tel,
		form.Password,
	)
	if err != nil {
		// 重复
		this.Data["Err"] = err.Error()
		this.TplName = "user/user_create.html"
		return
	}
	this.Redirect("/user", http.StatusFound)
}

type UserDeleteController struct {
	AuthController
}

func (this *UserDeleteController) Get() {
	id, _ := this.GetInt("id")
	service.Us.Delete(id)
	this.Redirect("/user", http.StatusFound)
}

type UserQueryController struct {
	AuthController
}

func (this *UserQueryController) Get() {
	s := this.GetString("q")
	userList := service.Us.Query(s)
	this.Data["Users"] = userList
	this.TplName = "user/user_list.html"
}

type UserModifyController struct {
	AuthController
}

func (this *UserModifyController) Get() {
	id, _ := this.GetInt("id")
	user, err := service.Us.Get(id)
	if err != nil {
		this.Redirect("/user", http.StatusFound)
		return
	}
	this.Data["User"] = user
	this.Data["Err"] = ""
	this.TplName = "user/user_modify.html"
}

func (this *UserModifyController) Post() {
	// id, _ := this.GetInt("id")
	// name := this.GetString("name")
	// brithday := this.GetString("brithday")
	// sex := this.GetString("sex")
	// addr := this.GetString("addr")
	// tel := this.GetString("tel")

	valid := validation.Validation{}
	form := forms.UserModifyForm{}
	this.ParseForm(&form)
	_, err := valid.Valid(&form)
	if err != nil {
		this.Data["Err"] = err.Error()
		this.Data["User"] = form
		this.TplName = "user/user_modify.html"
		return
	}

	// fmt.Println(form)
	err = service.Us.Modify(
		form.Id,
		form.Name,
		form.Brithday,
		form.Sex,
		form.Addr,
		form.Tel,
	)
	if err != nil {
		this.Data["Err"] = err.Error()
		this.Data["User"] = form
		this.TplName = "user/user_modify.html"
		return
	}
	this.Redirect("/user", http.StatusFound)
}

type UserPasswordController struct {
	AuthController
}

func (this *UserPasswordController) Get() {
	this.TplName = "user/changepassword.html"
}

func (this *UserPasswordController) Post() {
	// oldpassword := this.GetString("oldpassword")
	// newpassword := this.GetString("newpassword")
	// confirmpassword := this.GetString("confirmpassword")
	valid := validation.Validation{}
	form := forms.UserPasswordForm{}
	this.ParseForm(&form)
	user := this.User
	// bool?
	_, err := valid.Valid(&form)
	if err != nil {
		this.Data["Err"] = err.Error()
		this.TplName = "user/changepassword.html"
		return
	}
	if !user.CheckPassword(form.OldPassword) {
		this.Data["Err"] = "修改失败"
		this.TplName = "user/changepassword.html"
		return
	}
	err = service.Us.ChangePassword(user.Id, form.NewPassword)
	if err != nil {
		this.Data["Err"] = "修改失败"
		this.TplName = "user/changepassword.html"
		return
	}
	auth.Logout(this.Ctx)
	this.Redirect("/login", http.StatusFound)
}
