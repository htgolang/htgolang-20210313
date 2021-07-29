package utils

import (
	"cmdb/db"
	"database/sql"
	"fmt"
	"log"
)

// //检查用户名是否冲突
// func CheckUserName(name string, users []models.User) bool {
// 	for _, v := range users {
// 		if name == v.Name {
// 			return true
// 		}
// 	}
// 	return false
// }

// //检查号码是否冲突
// func CheckTel(tel string, users []models.User) bool {
// 	for _, v := range users {
// 		if tel == v.Telephone {
// 			return true
// 		}
// 	}
// 	return false
// }

func CheckUserExists(uid int64, field, value string) bool {
	var id int64 = -1
	var row *sql.Row

	if uid >= 0 {
		//如果是编辑用户，把当前编辑的用户排除在外
		sql := fmt.Sprintf("select id from user where %s=? and id != ?;", field)
		row = db.DbConn.QueryRow(sql, value, uid)
	} else {
		sql := fmt.Sprintf("select id from user where %s=?;", field)
		row = db.DbConn.QueryRow(sql, value)
	}

	if err := row.Scan(&id); err != nil {
		log.Printf("check_user_exists: scan data from row error.%s", err)
	}
	return id != -1
}
