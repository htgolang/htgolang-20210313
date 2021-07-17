package forms

import "github.com/beego/beego/v2/core/validation"

type UserCreateForm struct {
	Name     string `form:"name" valid:"Required"`
	Brithday string `form:"brithday"`
	Sex      string `form:"sex"`
	Addr     string `form:"addr"`
	Tel      string `form:"tel"`
	Password string `form:"password" valid:"Required;MinSize(6)"`
}

type UserModifyForm struct {
	Id       int    `form:"id" valid:"Required"`
	Name     string `form:"name" valid:"Required"`
	Brithday string `form:"brithday"`
	Sex      string `form:"sex"`
	Addr     string `form:"addr"`
	Tel      string `form:"tel"`
}

type UserPasswordForm struct {
	OldPassword     string `form:"oldpassword" valid:"Required"`
	NewPassword     string `form:"newpassword" valid:"Required"`
	ConfirmPassword string `form:"confirmpassword" valid:"Required"`
}

func (f *UserPasswordForm) Valid(v *validation.Validation) {
	if f.NewPassword != f.ConfirmPassword {
		v.SetError("ConfirmPassword", "两次输入的密码不一致")
	}
}
