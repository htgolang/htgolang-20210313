package models

import "time"

type RelRoleModule struct {
	RoleId   int64 `orm:"pk"`
	ModuleId int64
	CreateAt *time.Time `orm:"auto_now_add"`
	UpdateAt *time.Time `orm:"auto_now"`
	DeleteAt *time.Time `orm:"null"`
}
