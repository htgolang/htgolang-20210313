package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int64
	Name         string
	Password     string
	Birthday     *time.Time
	Telephone    string
	Email        string
	Addr         string
	Status       int8
	RoleId       int64
	DepartmentId int64
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
	Description  string
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
	return u.Birthday.Format("2006-01-02")
}
