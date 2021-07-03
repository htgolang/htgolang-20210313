package main

import (
	"fmt"
	"net/http"
	"usercmdb/config"
	_ "usercmdb/routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := ":9999"
	mysqlDriver := "mysql"
	dsn := "golang:golang@2021@tcp(192.168.0.158:3306)/cmdb?charset=utf8mb4&parseTime=true"
	err := config.OpenDb(mysqlDriver, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer config.CloseDB()

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
