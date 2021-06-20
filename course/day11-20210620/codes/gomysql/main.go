package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ddl(db *sql.DB) error {
	// 执行
	sql := `
	create table if not exists alarm(
		id bigint primary key auto_increment,
		alarm_time datetime not null,
		content text
	) engine=innodb default charset utf8mb4;
	`
	_, err := db.Exec(sql)
	return err
}

func dmlInsert(db *sql.DB) error {
	rs, err := db.Exec("insert into alarm(alarm_time, content) values('2021-6-20 17:05:10', 'CPU告警')")
	if err != nil {
		return err
	}
	fmt.Println(rs.LastInsertId())
	fmt.Println(rs.RowsAffected())

	alarmTime, alarmContent := time.Now().Format("2006-01-02 15:04:05"), "内存告警"

	// insert into alarm(t, c) values('', ''); delete from user;#')
	// sql注入
	sql := "insert into alarm(alarm_time, content) values('" + alarmTime + "', '" + alarmContent + "')"
	fmt.Println(sql)
	rs, err = db.Exec(sql)
	if err != nil {
		return err
	}
	fmt.Println(rs.LastInsertId())
	fmt.Println(rs.RowsAffected())

	// 预处理方式
	sql = "insert into alarm(alarm_time, content) values(?, ?);"
	alarmContent = "', ''); delete from user;#"
	rs, err = db.Exec(sql, alarmTime, alarmContent)
	if err != nil {
		return err
	}
	fmt.Println(rs.LastInsertId())
	fmt.Println(rs.RowsAffected())

	return nil
}

func dmlUpdate(db *sql.DB) error {
	sql := "update alarm set status=status+1 where id % 2=0"
	rs, err := db.Exec(sql)
	if err != nil {
		return err
	}
	fmt.Println(rs.LastInsertId())
	fmt.Println(rs.RowsAffected())

	sql = "update alarm set content=?, status=? where id=?"
	rs, err = db.Exec(sql, "update alarm content", 100, 1)
	if err != nil {
		return err
	}
	fmt.Println(rs.LastInsertId())
	fmt.Println(rs.RowsAffected())
	return nil

}

func dmlDelete(db *sql.DB) error {
	sql := "delete from alarm where id % 2=0 limit 1"
	rs, err := db.Exec(sql)
	if err != nil {
		return err
	}
	fmt.Println(rs.LastInsertId())
	fmt.Println(rs.RowsAffected())

	sql = "delete from alarm where id % 2 = ? limit ?"
	rs, err = db.Exec(sql, 1, 2)
	if err != nil {
		return err
	}
	fmt.Println(rs.LastInsertId())
	fmt.Println(rs.RowsAffected())
	return nil
}

func dqlRows(db *sql.DB) error {
	sql := "select id,content,alarm_time from alarm where content like ? limit ? offset ?"
	rows, err := db.Query(sql, "%内存%", 2, 5)
	if err != nil {
		return err
	}

	// 关闭rows 告知将连接放入连接池
	defer rows.Close()

	for rows.Next() {
		// 扫描查询数据到变量中，变量数量及类型需要与扫描结果中顺序和类型保持一致
		var id int64
		var content string
		var alarmTime time.Time
		if err := rows.Scan(&id, &content, &alarmTime); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(id, content, alarmTime)
		}
	}

	return nil
}

func dqlRow(db *sql.DB) error {
	sql := "select id, content, alarm_Time from alarm where id=?"
	var id int64
	var content string
	var alarmTime time.Time
	if err := db.QueryRow(sql, 51).Scan(&id, &content, &alarmTime); err != nil {
		return err
	}
	fmt.Println(id, content, alarmTime)
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

	fmt.Println("ok")
	fmt.Println("ddl:", ddl(db))

	// dml
	fmt.Println(dmlInsert(db))
	fmt.Println(dmlUpdate(db))
	fmt.Println(dmlDelete(db))

	// dql
	fmt.Println(dqlRows(db))
	fmt.Println(dqlRow(db))

}
