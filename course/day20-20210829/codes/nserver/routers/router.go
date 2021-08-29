package routers

import (
	"nserver/handlers"

	"github.com/gin-gonic/gin"
)

func Bind(engine *gin.Engine) {
	v1Group := engine.Group("/v1")
	{
		v1Group.POST("/register/", handlers.RegisterHandler)
		v1Group.POST("/heartbeat/", handlers.HeartbeatHandler)
		v1Group.GET("/config/", handlers.ConfigHandler)
		v1Group.GET("/tasks/", handlers.TasksHandler)
		v1Group.POST("/task/results/", handlers.ResultsHandler)
		v1Group.POST("/log/", handlers.LogHandler)
	}
}
