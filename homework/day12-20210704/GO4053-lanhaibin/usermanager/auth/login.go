package auth

import (
	context2 "context"
	"net/http"
	"usermanager/model"

	"github.com/beego/beego/v2/server/web/context"
)

func Login(ctx *context.Context, user model.User) {
	ctx.Input.CruSession.Set(context2.Background(), "uid", user.Id)
}

func Logout(ctx *context.Context) {
	ctx.Input.CruSession.Delete(context2.Background(), "uid")
}

func LoginRequired(ctx *context.Context) {
	uid := ctx.Input.Session("uid")
	if uid != nil {
		return
	}
	ctx.Redirect(http.StatusFound, "/login")
}
