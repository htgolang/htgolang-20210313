package models

import "time"

type User struct {
	Id        int `form:"id"`
	Name      string `form:"name"`
	Password  string `form:"password"`
	Phone     string `form:"phone"`
	Email     string `form:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	RoleId    int `form:"roleid"`
}

type ModifyTpl struct{
	Id        int
	Name      string
	Phone     string
	Email     string
	Role_name string
	SourceErr string
	NewErr    string
}

type Roles struct{
	Role_id int
	Role_name string
}