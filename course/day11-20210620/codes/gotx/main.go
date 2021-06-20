package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 打开数据库连接池
	dsn := "golang:golang@2021@tcp(10.0.0.2:3306)/cmdb?charset=utf8mb4&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// 只有在程序退出时关闭
	defer db.Close()

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	// 事务
	// 多个操作要同时成功 / 同时失败 多个操作放在一个事务中

	db.Exec("truncate table alarm;")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(tx)
	}

	sql := "insert into alarm(id, content, alarm_time) values(?, ?, ?)"
	_, err = tx.Exec(sql, 1, "第一条告警", time.Now())
	fmt.Println(err)
	if err == nil {
		_, err = tx.Exec(sql, 2, "第一条告警", time.Now())
		fmt.Println(err)
	}

	if err == nil {
		// 都成功 提交
		tx.Commit()
	} else {
		// 有一个失败 回退
		tx.Rollback()
	}
}
