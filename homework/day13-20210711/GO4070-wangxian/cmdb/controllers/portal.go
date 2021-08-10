package controllers

type PortalController struct {
	AuthenticationController
}

func (c *PortalController) Get() {
	c.TplName = "base/portal.html"
}
