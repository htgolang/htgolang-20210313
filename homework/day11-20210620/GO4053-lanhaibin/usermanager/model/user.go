package model

import (
	"crypto/sha512"
	"fmt"
	"strings"
	"time"
	"usermanager/db"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Brithday *time.Time
	Sex      bool
	Addr     string `json:"addr"`
	Tel      string `json:"tel"`
	Password string `json:"password"`
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

func (u User) Save() error {
	sql := "update user set name=?, brithday=?, sex=?, addr=?,tel=?,password=? where id=?"
	_, err := db.Db.Exec(sql,
		u.Name,
		u.Brithday,
		u.Sex,
		u.Addr,
		u.Tel,
		u.Password,
	)
	return err
}

func (u User) Create() (User, error) {
	sql := "insert into user (name, brithday, sex, addr, tel, password) values(?,?,?,?,?);"
	_, err := db.Db.Exec(sql,
		u.Name,
		u.Brithday,
		u.Sex,
		u.Addr,
		u.Tel,
		u.Password,
	)
	// 更新id
	return u, err
}

type UserMap map[int]*User

var Users UserMap = make(UserMap)

func GetUserId() int {
	var id int = 0
	for _, user := range Users {
		id = max(id, user.Id)
	}
	return id + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
