package forms

import (
	"cmdb/models"
	"cmdb/services"
	"strings"

	"github.com/beego/beego/v2/core/validation"
)

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"passwd"`
	User     *models.User
}

func (f *LoginForm) Valid(v *validation.Validation) {
	f.Username = strings.TrimSpace(f.Username)
	f.Password = strings.TrimSpace(f.Password)

	v.Required(f.Username, "username").Message("用户名不能为空")
	v.Required(f.Password, "password").Message("密码不能为空")

	if !v.HasErrors() {
		user := services.Login(f.Username, f.Password)
		if user == nil {
			v.SetError("user", "用户名密码错误")
		} else {
			f.User = user
		}
	}
}
