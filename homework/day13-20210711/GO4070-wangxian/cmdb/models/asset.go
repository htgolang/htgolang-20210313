package models

import "time"

type Asset struct {
	Id       int64      `orm:"pk"`
	Ip       string     `orm:"size(128)"`
	Addr     string     `orm:"size(128)"`
	CreateAt *time.Time `orm:"auto_now_add"`
	UpdateAt *time.Time `orm:"auto_now"`
	DeleteAt *time.Time `orm:"null"`
}
