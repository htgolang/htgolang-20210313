package models

import "time"

type Department struct {
	Id       int64
	Name     string
	CreateAt *time.Time
	UpdateAt *time.Time
	DeleteAt *time.Time
}
