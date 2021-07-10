package auth

import (
	"database/sql"
	"log"
	"webusermanage/config"
	"webusermanage/models"

	"golang.org/x/crypto/bcrypt"
)

func Auth(u, pwd string) *models.User{
	var users models.User
	db, err := sql.Open("mysql", config.DSN)
	if err != nil{
		log.Println("db err:",err)
	}
	defer db.Close()
	qsql := "select id,name,password,phone,email,role_id from user where name=?;"
	row, err := db.Query(qsql, u)
	if err != nil{
		log.Println("row err:",err)
	}
	defer row.Close()

	//校验密码
	for row.Next(){
		if err = row.Scan(&users.Id, &users.Name, &users.Password, &users.Phone, &users.Email, &users.RoleId); err == nil{
			if err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(pwd)); err == nil{
				return &users
			}	
		}
	}

	return nil
}