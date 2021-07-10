package forms

import (
	"fmt"
	"strings"
	"webusermanage/auth"
	"webusermanage/models"

	"github.com/beego/beego/v2/core/validation"
)

type Userform struct {
	Username string `form:"username"`
	Password string `form:"password"`

	User *models.User
}

func (f *Userform) Valid(v *validation.Validation) {
	f.Username = strings.TrimSpace(f.Username)
	f.Password = strings.TrimSpace(f.Password)

	v.Required(f.Username, "username").Message("用户名或密码不能为空")
	v.Required(f.Password, "password").Message("用户名或密码不能为空")
	fmt.Println(v.HasErrors())
	if !v.HasErrors() {
		if users := auth.Auth(f.Username, f.Password); users == nil {
			v.SetError("flag", "用户名或密码错误")
		} else {
			fmt.Println("uid:", users.Id)
			f.User = users
			fmt.Printf("1 User %#v\n:", f.User.Id)
		}
	}
}
