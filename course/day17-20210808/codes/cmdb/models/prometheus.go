package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type PrometheusAgent struct {
	Id       int64
	UUID     string `orm:"column(uuid);size(64);"`
	Addr     string `orm:"size(256);"`
	Hostname string `orm:"size(64)"`

	Config        string `orm:"type(text)"`
	ConfigVersion int64

	HeartbeatTime *time.Time

	CreatedAt *time.Time `orm:"auto_now_add;"`
	UpdatedAt *time.Time `orm:"auto_now;"`
	DeletedAt *time.Time `orm:"null;"`
}

func init() {
	orm.RegisterModel(new(PrometheusAgent))
}
