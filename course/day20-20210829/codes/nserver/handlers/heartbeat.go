package handlers

import (
	"net/http"

	"nserver/services"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func HeartbeatHandler(c *gin.Context) {
	// id
	logrus.Debug("hearbeat api")

	uuid, _ := c.GetQuery("uuid")
	services.OsqueryService.Heartbeat(uuid)
	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"error":  "",
		"result": nil,
	})
}
