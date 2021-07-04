package models

import "time"

type RelRoleModule struct {
	RoleId    int64
	ModuleId  int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
