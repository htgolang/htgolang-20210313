package models

import "time"

type User struct {
	ID        int64
	Name      string
	Sex       bool
	Addr      string
	Tel       string
	Birthday  string
	Passwd    string
	Create_at *time.Time
}

func NewUser(name string, sex bool, addr string, tel string, birthday string, passwd string) (*User, error) {
	createTime := time.Now()
	_, err := time.Parse("2006/01/02", birthday)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        -1,
		Name:      name,
		Sex:       sex,
		Addr:      addr,
		Tel:       tel,
		Birthday:  birthday,
		Passwd:    passwd,
		Create_at: &createTime,
	}, nil
}
