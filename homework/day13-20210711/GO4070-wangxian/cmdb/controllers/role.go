package controllers

import "cmdb/services"

type RoleController struct {
	AuthenticationController
}

func (c *RoleController) Query() {
	roles := services.RoleService.QueryRole()
	c.Data["Roles"] = roles
	c.TplName = "role/list.html"
}
