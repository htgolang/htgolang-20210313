package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Asset struct {
	Id          int64
	Addr        string     `orm:"size(256);"`
	Description string     `orm:"size(256);"`
	Status      int        // 0 未注册 1已注册
	CreatedAt   *time.Time `orm:"auto_now_add;"`
	UpdatedAt   *time.Time `orm:"auto_now;"`
	DeletedAt   *time.Time `orm:"null"`
}

func init() {
	orm.RegisterModel(new(Asset))
}
