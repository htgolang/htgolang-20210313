package controllers

type PortalController struct {
	LayoutController
}

func (c *PortalController) Get() {
	c.TplName = "base/portal.html"
}
