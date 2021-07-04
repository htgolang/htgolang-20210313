package services

import (
	"cmdb/models"
)

const sqlGetUserByName = "select id, name,password from user where name=? limit 1"

func ValidateLogin(name string, password string) *models.User {
	var user models.User
	err := db.QueryRow(sqlGetUserByName, name).Scan(&user.Id, &user.Name, &user.Password)
	// 先用明文，主要产品或项目中要使用bcrypt加密方式或者加盐sha256
	if err == nil && user.ValidPassword(password) {
		user.Password = ""
		return &user
	}
	return nil
}
