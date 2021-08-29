package routers

import (
	"cmdb/controllers"
	v1 "cmdb/controllers/apis/v1"

	"cmdb/filters"

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
	web.AutoRouter(new(controllers.OsqueryTaskController))

	web.AutoRouter(new(controllers.PrometheusAgentController))
	web.AutoRouter(new(controllers.PrometheusAlertController))

	web.Handler("/metrics/", promhttp.Handler())

	v1Namespace := web.NewNamespace("/v1",
		web.NSAutoRouter(new(v1.AgentController)),
		web.NSAutoRouter(new(v1.AlertController)),
	)

	web.AddNamespace(v1Namespace)
}
