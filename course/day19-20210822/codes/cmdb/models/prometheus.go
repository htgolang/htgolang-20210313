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

func (prometheusAgent *PrometheusAgent) IsHealthy() bool {
	if prometheusAgent.HeartbeatTime == nil {
		return false
	}

	return int64(time.Now().Sub(*prometheusAgent.HeartbeatTime)) < int64(time.Duration(5*time.Minute))
}

const (
	StatusAlertFiring   = 0
	StatusAlertResolved = 1
)

type PrometheusAlert struct {
	Id          int64
	Fingerprint string `orm:"size(32);"`
	Instance    string `orm:"size(256);"`
	Name        string `orm:"size(128);"`
	Description string `orm:"size(512);"`
	Severity    string `orm:"size(32)";`

	GeneratorURL string `orm:"column(generatorURL);size(1024);"`
	Labels       string `orm:"type(text);"`
	Annotations  string `orm:"type(text);"`

	Status   int // 0 触发 1解决
	StartsAt *time.Time
	EndsAt   *time.Time `orm:"null;"`

	CreatedAt *time.Time `orm:"auto_now_add;"`
	UpdatedAt *time.Time `orm:"auto_now;"`
	DeletedAt *time.Time `orm:"null;"`
}

func init() {
	orm.RegisterModel(new(PrometheusAgent))
	orm.RegisterModel(new(PrometheusAlert))
}
