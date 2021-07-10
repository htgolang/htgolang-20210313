package utils

import (
	"database/sql"
	"log"
	"webusermanage/models"

	"github.com/beego/beego/v2/server/web"
)

func GetUserById(id int) *models.User {
	var user models.User
	getusersql := "select id,name,role_id from user where id=?;"
	dsn, err := web.AppConfig.String("mysql::dsn")
	if err != nil {
		log.Fatalln(err)
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	row, err := db.Query(getusersql, id)
	if err != nil {
		log.Fatalln(err)
	}
	for row.Next() {
		if err = row.Scan(&user.Id, &user.Name, &user.RoleId); err != nil {
			return nil
		}
	}
	return &user
}
