package handlers

import (
	"encoding/json"
	"net/http"

	"nserver/services"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func TasksHandler(c *gin.Context) {
	// id
	logrus.Debug("get tasks api")

	uuid, _ := c.GetQuery("uuid")
	result, err := services.OsqueryService.Tasks(uuid)
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

func ResultsHandler(c *gin.Context) {
	// id body json
	logrus.Debug("task results api")
	uuid, _ := c.GetQuery("uuid")
	var result map[string]struct {
		Job    int64           `json:"job"`
		Status int             `json:"status"`
		Result json.RawMessage `json:"result"`
	}

	if err := c.BindJSON(&result); err == nil {
		for _, v := range result {
			services.OsqueryService.Result(uuid, v.Job, v.Status, v.Result)
		}
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"error":  "",
			"result": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":     false,
			"error":  err.Error(),
			"result": nil,
		})
	}

}
