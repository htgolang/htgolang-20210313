package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// 模型定义
type User struct {
	ID        int64      `orm:"column(id);pk;auto"`
	Name      string     `orm:"size(32)"`
	Password  string     `orm:"size(1024)"`
	Sex       bool       `orm:"-"`
	Tel       string     `orm:"size(32)"`
	Addr      string     `orm:"size(1024)"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
}

func NewUser(id int64, name string, sex bool, addr string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Password: "",
		Sex:      sex,
		Addr:     addr,
	}
}

// 注册数据库模型
func init() {
	orm.RegisterModel(new(User)) // &User{}
}
