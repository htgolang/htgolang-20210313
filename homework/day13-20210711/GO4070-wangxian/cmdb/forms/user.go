package forms

import (
	"cmdb/models"
	"cmdb/utils"
	"time"

	"github.com/beego/beego/v2/core/validation"
)

type AddUserForm struct {
	Username    string `form:"username"`
	Password    string `form:"password"`
	Confirmpw   string `form:"confirmpw"`
	Sex         int64  `form:"sex"`
	Birthday    string `form:"birthday"`
	Tel         string `form:"tel"`
	Email       string `form:"email"`
	Addr        string `form:"addr"`
	RoleId      int64  `form:"roleid"`
	Description string `form:"description"`
	User        *models.User
}

func (f *AddUserForm) Valid(v *validation.Validation) {
	v.Required(f.Username, "username.username").Message("用户名不能为空")
	v.Required(f.Password, "password.password").Message("密码不能为空")
	v.Required(f.Confirmpw, "confirmpw.confirmpw").Message("确认密码不能为空")
	v.Required(f.Birthday, "birthday.birthday").Message("生日不能为空")
	v.Required(f.Tel, "tel.tel").Message("手机号不能为空")
	v.Required(f.Email, "email.email").Message("邮箱不能为空")
	v.Required(f.Addr, "addr.addr.addr").Message("地址不能为空")
	v.MinSize(f.Password, 6, "password.password").Message("密码不少于6位")
	v.MaxSize(f.Password, 16, "password.password").Message("密码不超过16位")
	v.Mobile(f.Tel, "tel.tel").Message("手机号码错误")
	v.Email(f.Email, "email.email").Message("邮箱地址错误")

	birthday, err := time.Parse("2006-01-02", f.Birthday)
	if err != nil {
		v.SetError("birthday", "生日错误")
	}

	if f.Password != f.Confirmpw {
		v.SetError("confirmpw", "两次输入密码不一致!")
	}

	if utils.CheckUserExists(-1, "name", f.Username) {
		v.SetError("username", "用户名已存在")
	}

	if utils.CheckUserExists(-1, "telephone", f.Tel) {
		v.SetError("tel", "手机号已注册")
	}

	if utils.CheckUserExists(-1, "email", f.Email) {
		v.SetError("email", "邮箱已注册")
	}

	if !v.HasErrors() {
		f.User = new(models.User)
		f.User.Name = f.Username
		f.User.SetPassword(f.Password)
		f.User.Birthday = &birthday
		f.User.Sex = f.Sex == 1
		f.User.Telephone = f.Tel
		f.User.Email = f.Email
		f.User.Addr = f.Addr
		f.User.RoleId = f.RoleId
		f.User.Description = f.Description
	}
}

type EditUserForm struct {
	Id          int64  `form:"id"`
	Username    string `form:"username"`
	Sex         int64  `form:"sex"`
	Birthday    string `form:"birthday"`
	Tel         string `form:"tel"`
	Email       string `form:"email"`
	Addr        string `form:"addr"`
	RoleId      int64  `form:"roleid"`
	Description string `form:"description"`
	User        *models.User
}

func (f *EditUserForm) Valid(v *validation.Validation) {
	v.Required(f.Username, "username.username").Message("用户名不能为空")
	v.Required(f.Birthday, "birthday.birthday").Message("生日不能为空")
	v.Required(f.Tel, "tel.tel").Message("手机号不能为空")
	v.Required(f.Email, "email.email").Message("邮箱不能为空")
	v.Required(f.Addr, "addr.addr.addr").Message("地址不能为空")

	v.Mobile(f.Tel, "tel.tel").Message("手机号码错误")
	v.Email(f.Email, "email.email").Message("邮箱地址错误")

	birthday, err := time.Parse("2006-01-02", f.Birthday)
	if err != nil {
		v.SetError("birthday", "生日错误")
	}

	if utils.CheckUserExists(f.Id, "name", f.Username) {
		v.SetError("username", "用户名已存在")
	}

	if utils.CheckUserExists(f.Id, "telephone", f.Tel) {
		v.SetError("tel", "手机号已注册")
	}

	if utils.CheckUserExists(f.Id, "email", f.Email) {
		v.SetError("email", "邮箱已注册")
	}

	if !v.HasErrors() {
		f.User = new(models.User)
		f.User.Id = f.Id
		f.User.Name = f.Username
		f.User.Sex = (f.Sex == 1)
		f.User.Birthday = &birthday
		f.User.Telephone = f.Tel
		f.User.Email = f.Email
		f.User.Addr = f.Addr
		f.User.RoleId = f.RoleId
		f.User.Description = f.Description
	}
}

type ModifyPwForm struct {
	OldPassword string `form:"oldpassword"`
	NewPassword string `form:"newpassword"`
	ConfirmPw   string `form:"confirmpw"`
	CurrentUser *models.User
}

func (f *ModifyPwForm) Valid(v *validation.Validation) {
	v.Required(f.OldPassword, "oldpassword.oldpassword").Message("原密码不能为空")
	v.Required(f.NewPassword, "newpassword.newpassword").Message("新密码不能为空")
	v.Required(f.ConfirmPw, "confirmpw.confirmpw").Message("确认密码不能为空")
	v.MinSize(f.NewPassword, 6, "newpassword.newpassword").Message("新密码不少于6位")
	v.MaxSize(f.NewPassword, 16, "newpassword.newpassword").Message("新密码不超过16位")

	if !f.CurrentUser.ValidPassword(f.OldPassword) {
		v.SetError("oldpassword", "原密码错误")
		return
	}

	if f.NewPassword == f.OldPassword {
		v.SetError("newpassword", "新密码不能和原密码一致")
	}

	if f.NewPassword != f.ConfirmPw {
		v.SetError("confirmpw", "两次输入密码不一致")
	}

}
