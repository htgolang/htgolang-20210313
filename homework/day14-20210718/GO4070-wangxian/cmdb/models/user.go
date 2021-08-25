package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int64      `orm:"pk"`
	Name         string     `orm:"size(64)"`
	Password     string     `orm:"size(1024)"`
	Birthday     *time.Time `orm:"type(date);null"`
	Telephone    string     `orm:"size(64)"`
	Email        string     `orm:"size(64)"`
	Addr         string     `orm:"size(128)"`
	Status       int8
	RoleId       int64
	DepartmentId int64
	CreateAt     *time.Time `orm:"auto_now_add"`
	UpdateAt     *time.Time `orm:"auto_now"`
	DeleteAt     *time.Time `orm:"null"`
	Description  string     `orm:"null"`
	Sex          bool
}

func (u User) SetText() string {
	if u.Sex {
		return "男"
	}
	return "女"
}

func (u *User) SetPassword(password string) error {
	hashValue, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashValue)
	return nil
}

func (u *User) ValidPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
		return true
	}
	return false
}

func (u User) GetBirthday() string {
	if u.Birthday == nil {
		return ""
	}
	return u.Birthday.Format("2006-01-02")
}

func (u *User) TableName() string {
	return "user"
}
