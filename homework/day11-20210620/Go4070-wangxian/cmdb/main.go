package main

import (
	"cmdb/conf"
	"cmdb/db"
	_ "cmdb/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//初始化数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
	err := db.InitDB(conf.DbDriver, dsn)
	if err != nil {
		log.Fatalf("connect to db faild. %s", err)
	}
	defer db.DbConn.Close()

	http.ListenAndServe("172.24.239.212:9999", nil)
}
