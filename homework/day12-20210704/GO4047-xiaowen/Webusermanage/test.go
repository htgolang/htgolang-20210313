package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	pwd := "123"
	// pwd1 := "123"
	qsql := "select password from user;"
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	// hash1, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	fmt.Println(string(hash))
	fmt.Println(bcrypt.CompareHashAndPassword([]byte("$2a$10$dADVzJeOEfsBr7Map1jtRO3eQ5p.rCfSTQqbOW4suRQvCn1LAl6WO"), []byte("123")))
	fmt.Println(bcrypt.Cost(hash))




	fmt.Println("---------------------------------")
	dsn := "root:123.com@tcp(192.168.221.130:3306)/mycmdb?charset=utf8mb4&parseTime=true"
	db, _ := sql.Open("mysql", dsn)
	row, _ := db.Query(qsql)
	defer db.Close()
	var pwds string
	for row.Next(){
		row.Scan( &pwds)
		fmt.Println( pwds)
	}






	
	row.Close()

}