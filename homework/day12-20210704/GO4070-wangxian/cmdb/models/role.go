package models

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type Role struct {
	Id       int64
	Name     string
	Modules  string
	CreateAt *time.Time
	UpdateAt *time.Time
	DeleteAt *time.Time
}

func (r Role) GetModules() []string {
	var modules []string
	err := json.Unmarshal([]byte(r.Modules), &modules)
	if err != nil {
		log.Println(err)
		return nil
	}
	return modules
}

func (r Role) SetModulesText() string {
	rs := ""
	modules := r.GetModules()
	if modules == nil {
		return rs
	}
	for _, v := range modules {
		switch v {
		case "userManager":
			rs += "用户管理|"
		case "departmentManager":
			rs += "部门管理|"
		case "roleManager":
			rs += "角色管理|"
		case "assetManager":
			rs += "资产管理|"
		}
	}
	return strings.TrimRight(rs, "|")
}
