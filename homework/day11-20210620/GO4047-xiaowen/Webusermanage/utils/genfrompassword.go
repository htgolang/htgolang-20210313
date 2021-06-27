package utils

import (
	"database/sql"
	"log"
	"webusermanage/config"

	"golang.org/x/crypto/bcrypt"
)


func Genpwd(password string) (string,bool) {
	hashpwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		log.Println(err)
		return "nil",false
	}
	return string(hashpwd),true
}

func Comparepwd(hashpwd []byte, newpwd string)(string, bool){
	err := bcrypt.CompareHashAndPassword(hashpwd, []byte(newpwd))
	if err != nil{
		log.Println(err)
		return "nil",false
	}
	return Genpwd(newpwd)
}

func SourcePwdValidate(sourcepwd, name string)(bool){
	var pwd string
	db, err := sql.Open("mysql",config.DSN)
	if err != nil{
		log.Println(err)
	}
	defer db.Close()
	qsql := "select password from user where name=?"
	row, err := db.Query(qsql, name)
	if err != nil{
		log.Println(err)
	}
	defer row.Close()
	row.Next()
	row.Scan(&pwd)
	_, flags := Comparepwd([]byte(pwd), sourcepwd)
	if flags{
		return flags
	}
	return flags
}