package controllers

import "cmdb/services"

type LayoutController struct {
	AuthenticationController
}

func (c *LayoutController) Prepare() {
	c.AuthenticationController.Prepare()

	if c.CurrentUser == nil {
		return
	}

	currentTitle := "CMDB"
	currentMenu := ""
	currentModule := services.ModuleService.QueryModuleByUrl(c.Ctx.Input.URL())
	if currentModule != nil {
		currentTitle = currentModule.Name
		currentMenu = currentModule.KeyName
	}

	c.Layout = "base/layout.html"
	c.LayoutSections = map[string]string{}
	c.LayoutSections["CssSection"] = ""
	c.LayoutSections["JsSection"] = ""
	c.Data["CurrentTitle"] = currentTitle
	c.Data["CurrentMenu"] = currentMenu
}
