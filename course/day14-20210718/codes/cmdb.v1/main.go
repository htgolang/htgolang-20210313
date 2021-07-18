package main

import (
	_ "cmdb/routers"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn, err := web.AppConfig.String("mysql::dsn")
	if err != nil || dsn == "" {
		log.Fatal("error config dsn")
	}

	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err = orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		log.Fatal(err)
	}
	web.Run()
}
