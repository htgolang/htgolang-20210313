package main

import (
	"cmdb/config"
	_ "cmdb/routers"
	"cmdb/services"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 用户登陆逻辑
	// url => handler/handlerfunc => 业务逻辑 => template加载模板并输出

	// 定义Handler/HandlerFunc
	// 绑定关系 url handler
	// 启动服务

	err := services.InitDb(config.DbType, config.DbDSN)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(config.ServerAddr, nil))

}
