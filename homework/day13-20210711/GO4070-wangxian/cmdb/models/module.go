package models

import "time"

type Module struct {
	Id       int64  `orm:"pk"`
	Name     string `orm:"size(128)"`
	Url      string `orm:"size(256)"`
	IsMenu   bool
	CreateAt *time.Time `orm:"auto_now_add"`
	UpdateAt *time.Time `orm:"auto_now"`
	DeleteAt *time.Time `orm:"null"`
}
