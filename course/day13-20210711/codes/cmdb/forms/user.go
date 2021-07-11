package forms

import (
	"cmdb/models"
	"cmdb/services"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/validation"
)

type CreateUserForm struct {
	Name        string `form:"name"`
	Password    string `form:"password"`
	Password2   string `form:"password2"`
	Birthday    string `form:"birthday"`
	Telephone   string `form:"telephone"`
	Email       string `form:"email"`
	Description string `form:"description"`
	Sex         int    `form:"sex"`
	Addr        string `form:"addr"`

	User *models.User
}

func (f *CreateUserForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Password = strings.TrimSpace(f.Password)
	f.Password2 = strings.TrimSpace(f.Password2)
	f.Birthday = strings.TrimSpace(f.Birthday)
	f.Telephone = strings.TrimSpace(f.Telephone)
	f.Email = strings.TrimSpace(f.Email)
	f.Description = strings.TrimSpace(f.Description)
	f.Addr = strings.TrimSpace(f.Addr)

	// 校验

	// name 长度 [2, 32]
	// name 不能重复
	result := v.MinSize(f.Name, 2, "name.name.name").Message("用户名长度必须在2~32个字符之间")
	if result.Ok {
		result = v.MaxSize(f.Name, 32, "name.name.name").Message("用户名长度必须在2~32个字符之间")
	}
	if result.Ok {
		if user := services.UserService.GetByName(f.Name); user != nil {
			v.SetError("name", "用户名已存在")
		}
	}

	// password 长度 [6, 32]
	v.MinSize(f.Password, 6, "password.password.password").Message("密码长度必须在6~32个字符之间")
	v.MaxSize(f.Password, 32, "password.password.password").Message("密码长度必须在6~32个字符之间")
	if f.Password != f.Password2 {
		v.SetError("password2", "两次密码不一致")
	}

	birthday, err := time.Parse("2006-01-02", f.Birthday)
	if err != nil {
		v.SetError("birthday", "出生日志不正确")
	}
	v.Mobile(f.Telephone, "telephone.telephone.telephone").Message("手机号码格式不正确")
	v.Email(f.Email, "email.email.email").Message("Email格式不正确")

	v.MaxSize(f.Addr, 128, "addr.addr.addr").Message("地址长度必须在128个字符之内")
	v.MaxSize(f.Description, 1024, "description.description.description").Message("描述长度必须在1028个字符之内")

	if !v.HasErrors() {
		f.User = new(models.User)
		f.User.Name = f.Name
		f.User.SetPassword(f.Password)
		f.User.Birthday = &birthday
		f.User.Telephone = f.Telephone
		f.User.Email = f.Email
		f.User.Description = f.Description
		f.User.Sex = f.Sex == 1
	}
}

type ModifyPasswordForm struct {
	Password     string `form:"password"`
	NewPassword  string `form:"new_password"`
	NewPassword2 string `form:"new_password2"`

	CurrentUser *models.User
}

func (f *ModifyPasswordForm) Valid(v *validation.Validation) {
	f.Password = strings.TrimSpace(f.Password)
	f.NewPassword = strings.TrimSpace(f.NewPassword)
	f.NewPassword2 = strings.TrimSpace(f.NewPassword2)

	// password 长度 [6, 32]
	v.MinSize(f.Password, 6, "password.password.password").Message("密码长度必须在6~32个字符之间")
	v.MaxSize(f.Password, 32, "password.password.password").Message("密码长度必须在6~32个字符之间")

	v.MinSize(f.NewPassword, 6, "new_password.new_password.new_password").Message("密码长度必须在6~32个字符之间")
	v.MaxSize(f.NewPassword, 32, "new_password.new_password.new_password").Message("密码长度必须在6~32个字符之间")
	if f.NewPassword != f.NewPassword2 {
		v.SetError("new_password2", "两次密码不一致")
	}

	if !v.HasErrors() && f.CurrentUser != nil {
		// 验证旧密码
		if !f.CurrentUser.ValidPassword(f.Password) {
			v.SetError("password", "密码错误")
		}

	}
}
