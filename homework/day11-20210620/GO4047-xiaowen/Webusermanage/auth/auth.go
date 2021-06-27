package auth

import (
	"database/sql"
	"log"
	"webusermanage/config"
	"golang.org/x/crypto/bcrypt"
)

func Auth(u, pwd string) (string, bool){
	db, err := sql.Open("mysql", config.DSN)
	if err != nil{
		log.Println("db err:",err)
	}
	defer db.Close()
	qsql := "select id,password from user where name=?;"
	row, err := db.Query(qsql, u)
	if err != nil{
		log.Println("row err:",err)
	}
	defer row.Close()

	//校验密码
	for row.Next(){
		var id, password string
		if err = row.Scan(&id, &password); err == nil{
			if err = bcrypt.CompareHashAndPassword([]byte(password), []byte(pwd)); err == nil{
				return id, true
			}	
		}
	}

	return "",false
}