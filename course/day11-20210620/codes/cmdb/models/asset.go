package models

import "time"

type Asset struct {
	Id        int64
	Ip        string
	Addr      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
