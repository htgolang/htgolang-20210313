package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webusermanage/handlers"
	"webusermanage/services"

	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	AuthorizationController
}

func (c *UserController) Add() {
	if c.Ctx.Input.IsPost() {
		handlers.AddHande(&c.AuthorizationController.Controller)
		c.Redirect(web.URLFor("UserController.Query"), http.StatusFound)
		return
	}
	//显示添加页面
	c.TplName = "add/add.html"
}

func (c *UserController) Query() {
	web.ReadFromRequest(&c.Controller)
	users := handlers.QueryHande()
	c.TplName = "user/user.html"
	c.Data["Users"] = users
}

func (c *UserController) Delete() {
	id := c.GetString("id")
	services.Delete(id)
	c.Redirect(web.URLFor("UserController.Query"), http.StatusFound)
}

func (c *UserController) Modify() {
	id, aerr := strconv.Atoi(c.GetString("id"))
	if aerr != nil {
		log.Println(aerr)
	}
	handlers.ModifyHande(id, &c.Controller)
}

func (c *UserController) Datil() {
	id := c.GetString("id")
	flag, data := services.GetDtail(id)
	fmt.Fprintf(c.Ctx.ResponseWriter, "{\"ok\": %v, \"data\": %#v}", flag, data)
}

func (c *UserController) Role() {
	c.TplName = "role/role.html"
	c.Data["Temp"] = handlers.RoleHande()
}
