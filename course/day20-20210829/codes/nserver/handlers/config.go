package handlers

import (
	"net/http"

	"nserver/services"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func ConfigHandler(c *gin.Context) {
	// id // version
	logrus.Debug("get config api")
	uuid, _ := c.GetQuery("uuid")
	version, _ := c.GetQuery("version")
	result, err := services.OsqueryService.Config(uuid, version)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":     false,
			"error":  err.Error(),
			"result": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"error":  "",
			"result": result,
		})
	}

}
