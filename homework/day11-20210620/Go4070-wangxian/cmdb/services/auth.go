package services

import (
	"cmdb/db"
	"cmdb/models"
)

func Login(username, password string) *models.User {
	var user models.User
	sql := `select id,name,password from user where name=?;`

	err := db.DbConn.QueryRow(sql, username).Scan(&user.Id, &user.Name, &user.Password)
	if err == nil && user.ValidPassword(password) {
		return &user
	}
	return nil

}
