package main

import (
	"nserver/routers"
	"nserver/services"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	addr := ":10001"

	driverName := "mysql"
	dsn := "golang:golang@2021@tcp(10.0.0.2:3306)/cmdb?charset=utf8mb4&parseTime=true&loc=PRC"

	err := services.InitDb(driverName, dsn)
	if err != nil {
		logrus.Fatal(err)
	}
	router := gin.Default()
	routers.Bind(router)
	router.Run(addr)
}
