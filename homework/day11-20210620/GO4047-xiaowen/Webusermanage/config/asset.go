package config

import "time"

type Asset struct {
	Id        int
	Ip        string
	Addr      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
