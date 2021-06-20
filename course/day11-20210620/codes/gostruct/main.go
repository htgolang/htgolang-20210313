package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func toStruct(db *sql.DB, tableNames ...string) error {
	for _, tableName := range tableNames {
		sql := fmt.Sprintf("select * from %s limit 0", tableName)

		rows, err := db.Query(sql)
		if err != nil {
			return err
		}
		// 关闭rows 告知将连接放入连接池
		defer rows.Close()
		columns, err := rows.Columns()
		columnTypes, err := rows.ColumnTypes()

		// 处理错误
		genStruct(tableName, columns, columnTypes)
	}
	return nil
}

func snake(txt string) string {
	var builder strings.Builder
	var upper = true
	for _, ch := range txt {
		if ch == '_' {
			upper = true
			continue
		}
		if upper {
			builder.WriteString(strings.ToUpper(fmt.Sprintf("%c", ch)))
		} else {
			builder.WriteString(fmt.Sprintf("%c", ch))
		}
		upper = false
	}
	return builder.String()
}

func tlType(dbType string) string {
	m := map[string]string{
		"TINYINT":  "int8",
		"INT":      "int",
		"INTEGER":  "int",
		"BIGINT":   "int64",
		"CHAR":     "string",
		"VARCHAR":  "string",
		"TEXT":     "string",
		"DATE":     "*time.Time",
		"DATETIME": "*time.Time",
	}
	if typ, ok := m[dbType]; ok {
		return typ
	}
	return "string"
}

func genStruct(tableName string, columns []string, columnTypes []*sql.ColumnType) error {
	fhandler, err := os.Create(filepath.Join("objs", fmt.Sprintf("%s.go", strings.ToLower(tableName))))
	if err != nil {
		return err
	}
	defer fhandler.Close()

	fmt.Fprint(fhandler, "package objs\n\n")
	fmt.Fprintf(fhandler, "type %s struct {\n", snake(tableName))
	for i, typ := range columnTypes {
		fmt.Fprintf(fhandler, "\t %s %s\n", snake(columns[i]), tlType(typ.DatabaseTypeName()))
	}

	fmt.Fprintf(fhandler, "}")
	return nil
}

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

	rows, err := db.Query("show tables;")
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err == nil {
			toStruct(db, name)
		}
	}
	rows.Close()
}
