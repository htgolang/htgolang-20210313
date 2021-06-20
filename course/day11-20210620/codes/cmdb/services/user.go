package services

import (
	"cmdb/models"
	"database/sql"
	"fmt"
	"log"
)

const (
	sqlQueryUser       = "select id, name, sex from user"
	sqlQueryUserByName = "select id, name, sex from user where name like ?"
	sqlDeleteUserById  = "delete from user where id=? limit 1"
	sqlGetUserById     = "select id, name, sex from user where id=? limit 1"
)

func GetUserById(id int64) *models.User {
	var user models.User
	if err := db.QueryRow(sqlGetUserById, id).Scan(&user.Id, &user.Name, &user.Sex); err == nil {
		return &user
	}
	return nil
}

func QueryUsers(q string) []models.User {
	var rows *sql.Rows
	var err error
	if q == "" {
		rows, err = db.Query(sqlQueryUser)
	} else {
		rows, err = db.Query(sqlQueryUserByName, fmt.Sprintf("%%%s%%", q))
	}

	if err != nil {
		log.Println(err)
		return nil
	}
	var users = make([]models.User, 0)
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Sex); err != nil {
			log.Println(err)
			break
		}
		users = append(users, user)
	}
	return users
}

func DeleteUser(id int64) {
	db.Exec(sqlDeleteUserById, id)
}
