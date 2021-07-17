package model

import (
	"crypto/sha512"
	"fmt"
	"strings"
	"time"
)

type User struct {
	Id       int
	Name     string `orm:"size(64)"`
	Brithday *time.Time
	Sex      bool
	Addr     string `orm:"size(255)"`
	Tel      string `orm:"size(64)"`
	Password string `orm:"size(255)"`
}

func (u User) String() string {
	return fmt.Sprintf("id: %d, name: %s, addr: %s, tel: %s",
		u.Id, u.Name, u.Addr, u.Tel)
}

func (u User) Contains(s string) bool {
	if strings.Contains(u.Name, s) {
		return true
	}
	if strings.Contains(u.Addr, s) {
		return true
	}
	if strings.Contains(u.Tel, s) {
		return true
	}
	return false
}

func (u User) CheckPassword(password string) bool {
	hash := sha512.New()
	hash.Write([]byte(password))
	return u.Password == fmt.Sprintf("%x", string(hash.Sum([]byte{})))
}

func (u *User) SetPassword(password string) {
	hash := sha512.New()
	hash.Write([]byte(password))
	u.Password = fmt.Sprintf("%x", hash.Sum(nil))
}

// func (u User) Save() error {
// 	sql := "update user set name=?, brithday=?, sex=?, addr=?,tel=?,password=? where id=?"
// 	_, err := db.Db.Exec(sql,
// 		u.Name,
// 		u.Brithday,
// 		u.Sex,
// 		u.Addr,
// 		u.Tel,
// 		u.Password,
// 	)
// 	return err
// }

// func (u User) Create() (User, error) {
// 	sql := "insert into user (name, brithday, sex, addr, tel, password) values(?,?,?,?,?);"
// 	_, err := db.Db.Exec(sql,
// 		u.Name,
// 		u.Brithday,
// 		u.Sex,
// 		u.Addr,
// 		u.Tel,
// 		u.Password,
// 	)
// 	return u, err
// }
