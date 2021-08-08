package routers

import (
	"cmdb/controllers"
	"cmdb/filters"

	"cmdb/controllers/apis/v1"

	"github.com/beego/beego/v2/server/web"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	web.ErrorController(new(controllers.ErrorController))

	web.InsertFilter("/*", web.BeforeExec, filters.BeforeExec)
	web.InsertFilter("/*", web.AfterExec, filters.AfterExec, web.WithReturnOnOutput(false))

	web.Router("/", new(controllers.HomeController), "*:Index")
	web.AutoRouter(new(controllers.AuthController))
	web.AutoRouter(new(controllers.UserController))
	web.AutoRouter(new(controllers.AssetController))

	web.Handler("/metrics/", promhttp.Handler())

	v1Namespace := web.NewNamespace("/v1",
		web.NSAutoRouter(new(v1.AgentController)),
	)

	web.AddNamespace(v1Namespace)
}
