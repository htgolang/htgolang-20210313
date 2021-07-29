package utils

import (
	"cmdb/db"
	"database/sql"
	"log"
)

func CheckDepartmentExists(dpid int64, name string) bool {
	var row *sql.Row
	var id int64 = -1
	// fmt.Println(dpid, name)
	if dpid >= 0 {
		// fmt.Println("----------")
		sql := `select id from department where name=? and id != ?;`
		row = db.DbConn.QueryRow(sql, name, dpid)
	} else {
		sql := `select id from department where name=?;`
		row = db.DbConn.QueryRow(sql, name)
	}
	err := row.Scan(&id)
	if err != nil {
		log.Printf("check dp exists: scan data from row error. %s", err)
	}
	// fmt.Println(id)

	return id != -1
}
