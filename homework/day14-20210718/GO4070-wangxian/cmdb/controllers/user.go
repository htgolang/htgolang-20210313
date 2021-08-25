package controllers

import (
	"cmdb/forms"
	"cmdb/models"
	"cmdb/services"
	"log"
	"net/http"

	"github.com/beego/beego/v2/adapter/validation"
)

type UserController struct {
	LayoutController
}

func (c *UserController) Query() {

	keyword := c.GetString("q")
	users := services.UserService.QueryUser(keyword)
	// fmt.Println(users)
	c.Data["Users"] = users
	c.TplName = "user/list.html"

}

func (c *UserController) Delete() {
	id, _ := c.GetInt64("id")
	services.UserService.DeleteUser(id)
	c.Redirect("/user/query/", http.StatusFound)
}

func (c *UserController) Add() {
	var form forms.AddUserForm
	var valid validation.Validation
	roles := services.RoleService.QueryRole()

	if c.Ctx.Input.IsPost() {
		err := c.ParseForm(&form)
		// fmt.Println(form)
		if err == nil {
			if b, err := valid.Valid(&form); err != nil {
				log.Println("valid add user form data error.", err)
				valid.SetError("user", "验证数据错误")
			} else if b {
				err := services.UserService.AddUser(form.User)
				if err == nil {
					c.Redirect("/user/query/", http.StatusFound)
					return
				}
				log.Println("add user faild.", err)
				valid.SetError("user", "添加用户失败")
			}

		} else {
			// fmt.Println(err)
			valid.SetError("user", "提交数据错误")
		}
	}

	// fmt.Println(valid.ErrorsMap)
	c.Data["Roles"] = roles
	c.Data["ErrMsgs"] = valid.ErrorMap()
	c.TplName = "user/add.html"
}

func (c *UserController) Edit() {
	var form forms.EditUserForm
	var valid validation.Validation
	var user *models.User = &models.User{}
	roles := services.RoleService.QueryRole()

	if c.Ctx.Input.IsGet() {
		id, _ := c.GetInt64("id")
		user = services.UserService.QueryUserByID(id)
		form.Model2From(user)
	}

	if c.Ctx.Input.IsPost() {
		err := c.ParseForm(&form)
		// fmt.Printf("%#v", form)
		if err == nil {
			if b, err := valid.Valid(&form); err != nil {
				log.Println("valid edit user form data error.", err)
				valid.SetError("user", "验证数据错误")
			} else if b {
				//更新数据
				err := services.UserService.EditUser(form.User)
				if err == nil {
					c.Redirect("/user/query/", http.StatusFound)
					return
				}
				log.Println("edit user info faild.", err)
				valid.SetError("user", "修改用户信息失败")
			}
		} else {
			valid.SetError("user", "提交数据错误")
		}
	}

	c.Data["Form"] = form
	c.Data["Roles"] = roles
	c.Data["ErrMsgs"] = valid.ErrorMap()
	c.TplName = "user/edit.html"
}

func (c *UserController) ModifyPw() {
	var valid validation.Validation
	var form forms.ModifyPwForm
	var success string

	if c.Ctx.Input.IsPost() {
		//获取当前用户
		form.CurrentUser = c.CurrentUser

		err := c.ParseForm(&form)

		if err == nil {
			if b, err := valid.Valid(&form); err != nil {
				log.Println("valid modifypw data error.", err)
				valid.SetError("userpassword", "验证数据错误")
			} else if b {
				form.CurrentUser.SetPassword(form.NewPassword)
				err := services.UserService.ModifyPw(form.CurrentUser)
				if err == nil {
					//修改密码成功，销毁session,重新登录
					c.DestroySession()
					success = "修改密码成功，请刷新页面重新登录"
				} else {
					//修改失败
					log.Println("modify password error.", err)
					valid.SetError("userpassword", "修改密码失败")
				}
			}
		} else {
			log.Println("parse modifypw data error.", err)
			valid.SetError("userpassword", "提交数据错误")
		}
	}

	c.Data["Success"] = success
	c.Data["ErrMsgs"] = valid.ErrorMap()
	c.TplName = "user/changepw.html"
}
