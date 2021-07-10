package services

import (
	"database/sql"
	"log"
	"strconv"
	"time"
	"webusermanage/config"
	"webusermanage/models"

	"github.com/beego/beego/v2/server/web"
)

func Add(user *models.User) (map[string]bool, error) {
	isql := "insert into user(name, password, phone, email, created_at, updated_at, role_id) values(?, ?, ?, ?, ?, ?, ?)"
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	_, err = db.Exec(isql, user.Name, user.Password, user.Phone, user.Email, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), user.RoleId)
	if err != nil {
		log.Println(err)
	}
	return map[string]bool{"ok": true}, nil
}

//3 用户查询
func Query(id string) models.User {
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	var i int
	var created_at, updated_at time.Time
	var name, phone, email string
	qsql := "select id, name, phone, email, created_at, updated_at from user where id=?;"
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	row, err := db.Query(qsql, num)
	if err != nil {
		log.Println(err)
	}
	defer row.Close()
	for row.Next() {
		if err = row.Scan(&i, &name, &phone, &email, &created_at, &updated_at); err == nil {
			return models.User{
				Id:        i,
				Name:      name,
				Phone:     phone,
				Email:     email,
				CreatedAt: created_at,
				UpdatedAt: updated_at,
			}
		}

	}
	return models.User{}
}

func QueryAll() []models.User {
	u := models.User{}
	us := make([]models.User, 0)
	qsql := "select id,name,email,phone,password,role_id from user;"
	dsn, _ := web.AppConfig.String("mysql::dsn")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	row, err := db.Query(qsql)
	if err != nil {
		log.Fatalln(err)
	}
	for row.Next() {
		var name, email, phone, password string
		var id int
		var role_id int
		row.Scan(&id, &name, &email, &phone, &password, &role_id)
		u.Id = id
		u.Name = name
		u.Email = email
		u.Password = password
		u.Phone = phone
		u.RoleId = role_id
		us = append(us, u)
	}
	defer row.Close()
	return us
}

//4 用户删除
func Delete(id string) map[string]bool {
	dsql := "delete from user where id=?"
	n, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	dsn, _ := web.AppConfig.String("mysql::dsn")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	_, err = db.Exec(dsql, n)
	if err != nil {
		log.Println(err)
		return map[string]bool{"ok": false}
	}

	return map[string]bool{"ok": true}
}

//5 用户信息修改
func Modif(user *models.User, id string) map[string]bool {
	usql := "update user set password=?, phone=?, email=?, updated_at=?;"
	dsn, _ := web.AppConfig.String("mysql::dsn")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	_, err = db.Exec(usql, user.Password, user.Phone, user.Email, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println(err)
		return map[string]bool{"ok": false}
	}
	return map[string]bool{"ok": true}
}

//获取详情
func GetDtail(id string) (bool, *models.User) {
	var users  models.User
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	qsql := "select id, name, phone, email, created_at, updated_at, role_id from user where id=?;"
	dsn,_ := web.AppConfig.String("mysql::dsn")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	row, err := db.Query(qsql, num)
	if err != nil {
		log.Println(err)
	}
	defer row.Close()
	for row.Next() {
		if err = row.Scan(&users.Id, &users.Name, &users.Phone, &users.Email, &users.CreatedAt, &users.UpdatedAt, &users.RoleId); err == nil {
			return true, &users
		}

	}
	return false, nil
}
