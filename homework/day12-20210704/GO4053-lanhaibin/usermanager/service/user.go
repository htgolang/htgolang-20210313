package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"usermanager/db"
	"usermanager/model"
)

const (
	sqlQueryUserById      = "select id, name, brithday, sex, addr, tel, password from user where id=?;"
	sqlQueryUserByName    = "select id, name, brithday, sex, addr, tel, password from user where name like ?;"
	sqlQueryUser          = "select id, name, brithday, sex, addr, tel, password from user;"
	sqlDeleteUserById     = "delete from user where id=?;"
	sqlUpdateUserById     = "update user set name=?, brithday=?, sex=?,addr=?,tel=? where id=?;"
	sqlUpdateUserPassword = "update user set password=? where id=?;"
)

type UserService struct{}

func (u UserService) Get(id int) (model.User, error) {
	user := model.User{}
	rows, err := db.Db.Query(sqlQueryUserById, id)
	if err == nil && rows.Next() {
		var brithday sql.NullString
		err = rows.Scan(&user.Id, &user.Name, &brithday, &user.Sex, &user.Addr, &user.Tel, &user.Password)
		if brithday.Valid {
			t, err := time.Parse("2006-02-01 15:04:05", brithday.String)
			if err == nil {
				user.Brithday = &t
			}
		}
		return user, nil
	}
	return user, errors.New("用户不存在!")
}

func (u UserService) Create(name, brithday string, sex bool, addr, tel, password string) error {
	sql := "insert into user (name, brithday, sex, addr, tel, password) values(?,?,?,?,?,?);"
	if len(name) == 0 || len(password) == 0 {
		return errors.New("用户名和密码不能为空")
	}
	user, err := u.GetUserByName(name)
	if err == nil {
		return errors.New("用户已存在")
	}
	user = model.User{
		Name:     name,
		Brithday: nil,
		Sex:      sex,
		Addr:     addr,
		Tel:      tel,
	}
	user.SetPassword(password)
	_, err = db.Db.Exec(sql, name, brithday, sex, addr, tel, user.Password)
	return err
}

func (u UserService) Delete(id int) error {
	_, err := db.Db.Exec(sqlDeleteUserById, id)
	return err
}

func (u UserService) Query(s string) (qs []model.User) {
	var rows *sql.Rows
	var err error
	if len(s) != 0 {
		s = fmt.Sprintf("%%%s%%", s)
		rows, err = db.Db.Query(sqlQueryUserByName, s)
	} else {
		sql := sqlQueryUser
		rows, err = db.Db.Query(sql)
	}
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		user := model.User{}
		var brithday sql.NullString
		err := rows.Scan(&user.Id, &user.Name, &brithday, &user.Sex, &user.Addr, &user.Tel, &user.Password)
		if err != nil {
			continue
		}
		if brithday.Valid {
			t, err := time.Parse("2006-02-01 15:04:05", brithday.String)
			if err == nil {
				user.Brithday = &t
			}
		}
		qs = append(qs, user)
	}
	return
}

func (u UserService) Modify(id int, name, brithday, sex, addr, tel string) error {
	user, err := u.GetUserById(id)
	if err != nil {
		return fmt.Errorf("user id %d not exists", id)
	}
	if len(name) > 0 {
		u, err := u.GetUserByName(name)
		if err == nil && u.Id != id {
			return errors.New("用户名已存在")
		}
		user.Name = name
	}
	_, err = db.Db.Exec(sqlUpdateUserById, name, brithday, sex, addr, tel, id)
	return err
}

func (u UserService) GetUserByName(name string) (model.User, error) {
	user := model.User{}
	rows, err := db.Db.Query(sqlQueryUserByName, name)
	if err == nil && rows.Next() {
		var brithday sql.NullString
		err := rows.Scan(&user.Id, &user.Name, &brithday, &user.Sex, &user.Addr, &user.Tel, &user.Password)
		if brithday.Valid {
			t, _ := time.Parse("2006-01-02 15:04:05", brithday.String)
			user.Brithday = &t
		}
		return user, err
	}
	return user, errors.New("用户不存在")
}

func (u UserService) GetUserById(id int) (model.User, error) {
	user := model.User{}
	rows, err := db.Db.Query(sqlQueryUserById, id)
	if err != nil {
		return user, err
	}
	if rows.Next() {
		var brithday sql.NullString
		err = rows.Scan(&user.Id, &user.Name, &brithday, &user.Sex, &user.Addr, &user.Tel, &user.Password)
		if brithday.Valid {
			t, _ := time.Parse("2006-01-02 15:04:05", brithday.String)
			user.Brithday = &t
		}
	}
	return user, err
}

func (u UserService) ChangePassword(id int, password string) error {
	user, _ := u.GetUserById(id)
	user.SetPassword(password)
	_, err := db.Db.Exec(sqlUpdateUserPassword, user.Password, id)
	return err

}

var Us UserService
