package utils

import (
	"cmdb/db"
	"database/sql"
	"log"
)

func CheckAssetExists(id int64, ip string) bool {
	var aid int64 = -1
	var row *sql.Row
	if id >= 0 {
		sql := `select id from asset where ip=? and id!=?;`
		row = db.DbConn.QueryRow(sql, ip, id)
	} else {
		sql := `select id from asset where ip=?;`
		row = db.DbConn.QueryRow(sql, ip)
	}

	if err := row.Scan(&aid); err != nil {
		log.Printf("check_asset_exists: scan data from row error. %s", err)
	}
	// fmt.Println(aid)

	return aid != -1

}
