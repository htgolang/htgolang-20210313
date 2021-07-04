package main

import (
	_ "cmdb/routers"
	"cmdb/services"
	"log"

	"github.com/beego/beego/v2/server/web"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn, err := web.AppConfig.String("mysql::dsn")
	if err != nil || dsn == "" {
		log.Fatal("error config dsn")
	}
	err = services.InitDb("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	web.Run()
}
