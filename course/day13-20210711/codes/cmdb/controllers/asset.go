package controllers

type AssetController struct {
	AuthorizationController
}

func (c *AssetController) Query() {
	c.TplName = "asset/query.html"
}
