package auth

import (
	"usermanager/model"
	"usermanager/service"

	"github.com/beego/beego/v2/server/web/context"
)

func GetCurrentUser(ctx *context.Context) model.User {
	id := ctx.Input.Session("uid")
	if id == nil {
		return model.User{}
	}
	user, _ := service.Us.GetUserById(id.(int))
	return user
}
