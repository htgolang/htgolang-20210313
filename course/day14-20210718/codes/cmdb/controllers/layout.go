package controllers

import "cmdb/services"

type LayoutController struct {
	AuthorizationController
}

func (c *LayoutController) Prepare() {
	c.AuthorizationController.Prepare()

	if c.CurrentUser == nil {
		return
	}

	controllerName, actionName := c.GetControllerAndAction()

	currentTitle, currentMenu := "CMDB", ""
	if module := services.RoleService.GetModule(controllerName, actionName); module != nil {
		currentTitle = module.Name
		currentMenu = module.Key
	}

	c.Layout = "base/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["SectionStyles"] = ""
	c.LayoutSections["SectionScripts"] = ""
	c.Data["CurrentTitle"] = currentTitle
	c.Data["CurrentMenu"] = currentMenu
}
