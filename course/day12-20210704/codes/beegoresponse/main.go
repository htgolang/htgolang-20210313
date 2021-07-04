package main

import (
	"github.com/beego/beego/v2/server/web"
)

type ResponseController struct {
	web.Controller
}

func (c *ResponseController) Index() {
	// Response ResponseWriter
	// Ouput => Body Download
	// Ctx WriteString, Redirect
	// controller => TplName + Data
	// controller => Data + ServeJSON/SerXML
	c.Ctx.Output.Body([]byte("index"))
}

func (c *ResponseController) Download() {
	// Response ResponseWriter
	// Ouput => Body Download
	// Ctx WriteString, Redirect
	// controller => TplName + Data
	// controller => Data + ServeJSON/SerXML
	c.Ctx.Output.Download("./go.mod")
}

func (c *ResponseController) Json() {
	c.Data["json"] = map[string]string{"name": "kk", "addr": "西安"}
	c.ServeJSON()
}

func main() {
	web.AutoRouter(new(ResponseController))
	web.Run()
}
