package config

import "time"

type User struct {
	Id        int
	Name      string
	Password  string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	RoleId    int
}
