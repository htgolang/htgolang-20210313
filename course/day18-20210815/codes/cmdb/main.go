package main

import (
	_ "cmdb/routers"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"

	_ "cmdb/templates"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/beego/beego/v2/server/web/session/redis"
)

func main() {
	dsn, err := web.AppConfig.String("mysql::dsn")
	if err != nil || dsn == "" {
		log.Fatal("error config dsn")
	}

	web.BConfig.Log.AccessLogs = true
	logs.SetLogger(logs.AdapterFile, `{
		"filename":"logs/web.log",
		"level":7,
		"maxsize":10485760,
		"daily":true,
		"maxdays":10
	}`)

	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err = orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		log.Fatal(err)
	}

	orm.RunSyncdb("default", false, true)
	web.Run()
}
