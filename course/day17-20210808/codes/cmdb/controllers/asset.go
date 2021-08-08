package controllers

type AssetController struct {
	LayoutController
}

func (c *AssetController) Query() {
	c.TplName = "asset/query.html"
}
