package forms

import (
	"cmdb/models"
	"cmdb/services"
	"strings"

	"github.com/beego/beego/v2/core/validation"
)

type LoginForm struct {
	Username string `form:"name"`
	Password string `form:"password"`

	User *models.User
}

func (f *LoginForm) Valid(v *validation.Validation) {
	f.Username = strings.TrimSpace(f.Username)
	f.Password = strings.TrimSpace(f.Password)

	v.Required(f.Username, "username").Message("用户名或密码不能为空")
	v.Required(f.Password, "password").Message("用户名或密码不能为空")

	if !v.HasErrors() {
		if user := services.UserService.ValidateLogin(f.Username, f.Password); user == nil {
			v.SetError("user", "用户名或密码错误")
		} else {
			f.User = user
		}
	}
}
