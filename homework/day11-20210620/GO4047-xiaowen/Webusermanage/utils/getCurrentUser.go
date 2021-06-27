package utils

import (
	"database/sql"
	"log"
	"webusermanage/config"
)

func GetCurrentUser(uid string) *config.UserStructure{
	var user config.UserStructure
	qsql := "select id,name,role_id from user where id=?"
	db, err := sql.Open("mysql",config.DSN)
	if err != nil{
		log.Println(err)
		return nil
	}
	defer db.Close()
	row, err := db.Query(qsql, uid)
	if err != nil{
		log.Println(err)
		return nil
	}
	defer row.Close()
	row.Next()
	row.Scan(&user.Id, &user.UserName, &user.Role_id)
	return &user
}