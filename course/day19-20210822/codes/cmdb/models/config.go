package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

/*
key: scan
value: {""}
*/

type Config struct {
	Id        int64
	Key       string     `orm:"size(256);"`
	Value     string     `orm:"type(text);"`
	CreatedAt *time.Time `orm:"auto_now_add;"`
	UpdatedAt *time.Time `orm:"auto_now;"`
	DeletedAt *time.Time `orm:"null;"`
}

func init() {
	orm.RegisterModel(new(Config))
}
