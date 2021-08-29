package handlers

import (
	"net/http"

	"nserver/services"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	// id body json
	logrus.Debug("register api")
	var req struct {
		UUID     string `json:"uuid"`
		Hostname string `json:"hostname"`
		Addr     string `json:"addr"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":     false,
			"error":  err.Error(),
			"result": nil,
		})
	} else {
		services.OsqueryService.Register(req.UUID, req.Hostname, req.Addr)
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"error":  "",
			"result": nil,
		})
	}

}
