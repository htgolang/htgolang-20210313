package main

import (
	"database/sql"
	"fmt"
	"log"

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

	fmt.Println("ok")
	sql := "select * from user limit 0"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}

	// 关闭rows 告知将连接放入连接池
	defer rows.Close()

	columnTypes, _ := rows.ColumnTypes()
	columns, _ := rows.Columns()
	for i, typ := range columnTypes {
		fmt.Println(i, columns[i], "=======")
		fmt.Println(typ.DatabaseTypeName())
		fmt.Println(typ.Length())
		fmt.Println(typ.Nullable())
		fmt.Println(typ.ScanType())
		fmt.Println(typ.DecimalSize())
	}

}
