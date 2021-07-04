package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

// 定义控制器 http.Handler
type IndexController struct {
	web.Controller
}

func (c *IndexController) Get() { // ServeHTTP
	// 处理Get请求
	// c.Ctx.WriteString("hi kk") // fmt.Fprint(w, "hi kk")
	c.Data["user"] = map[string]string{"id": "1", "name": "kk"} // 给模板传递数据
	c.TplName = "index.html"                                    // 显示index.html 页面
}

func (c *IndexController) Post() {
	c.Ctx.WriteString("post")
}
