package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root@(127.0.0.1)/usermanager"

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}
