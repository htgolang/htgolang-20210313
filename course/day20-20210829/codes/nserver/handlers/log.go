package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func LogHandler(c *gin.Context) {
	// id  body json
	logrus.Debug("log api")
	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"error":  "",
		"result": nil,
	})
}
