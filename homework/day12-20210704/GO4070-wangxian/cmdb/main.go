package main

import (
	"cmdb/db"
	_ "cmdb/routers"
	"fmt"
	"log"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	dbuser, _ := web.AppConfig.String("db::DbUser")
	dbpassword, _ := web.AppConfig.String("db::DbPassword")
	dbhost, _ := web.AppConfig.String("db::DbHost")
	dbport, _ := web.AppConfig.String("db::DbPort")
	dbname, _ := web.AppConfig.String("db::DbName")
	dbtype, _ := web.AppConfig.String("db::DbType")

	//初始化数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", dbuser, dbpassword, dbhost, dbport, dbname)
	err := db.InitDB(dbtype, dsn)
	if err != nil {
		log.Fatalf("connect to db faild. %s", err)
	}
	defer db.DbConn.Close()

	web.Run()
}
