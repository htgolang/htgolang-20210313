package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int64
	Name         string     `orm:"size(64);"`
	Password     string     `orm:"size(1024);"`
	Birthday     *time.Time `orm:"type(date);null"`
	Telephone    string     `orm:"size(64);"`
	Email        string     `orm:"size(64);"`
	Addr         string     `orm:"size(128);"`
	Status       int8
	RoleId       int64
	DepartmentId int64
	CreatedAt    *time.Time `orm:"auto_now_add;"`
	UpdatedAt    *time.Time `orm:"auto_now;"`
	DeletedAt    *time.Time `orm:"null"`
	Description  string     `orm:type(text)`
	Sex          bool
}

func (u *User) SexText() string {
	if u.Sex {
		return "男"
	}
	return "女"
}

func (u *User) SetPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func (u *User) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func init() {
	orm.RegisterModel(new(User))
}
