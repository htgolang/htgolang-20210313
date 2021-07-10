package utils

import (
	"database/sql"
	"log"
	"webusermanage/models"

	"github.com/beego/beego/v2/server/web"
)

func GetCurrentUser(uid int) *models.User{
	var user models.User
	qsql := "select id,name,role_id from user where id=?"
	dsn,_ := web.AppConfig.String("mysql::dsn")
	db, err := sql.Open("mysql",dsn)
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
	row.Scan(&user.Id, &user.Name, &user.RoleId)
	return &user
}