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

	v.Required(f.Username, "user.user").Message("用户名或密码不能为空")
	v.Required(f.Password, "user.user").Message("用户名或密码不能为空")

	if !v.HasErrors() {
		user := services.UserService.QueryUserByName(f.Username)
		if user != nil && user.ValidPassword(f.Password) {
			f.User = user
		} else {
			v.SetError("user", "用户名密码错误")
		}
	}
}
