package models

import "time"

type Asset struct {
	Id       int64
	Ip       string
	Addr     string
	CreateAt *time.Time
	UpdateAt *time.Time
	DeleteAt *time.Time
}
