package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type OsqueryAgent struct {
	Id            int64
	UUID          string `orm:"column(uuid);size(64);"`
	Hostname      string `orm:"size(64);"`
	Addr          string `orm:"size(256);"`
	Config        string `orm:"type(text);"`
	ConfigVersion string `orm:"size(64);"`

	HeartbeatTime *time.Time

	CreatedAt *time.Time `orm:"auto_now_add;"`
	UpdatedAt *time.Time `orm:"auto_now;"`
	DeletedAt *time.Time `orm:"null;"`
}

type OsqueryTaskType struct {
	Id   int64
	Key  string `orm:"size(64);"`
	Name string `orm:"size(64);"`
	SQL  string `orm:"column(sql);type(text);"`

	Columns   [][2]string `orm:"-"`
	CreatedAt *time.Time  `orm:"auto_now_add;"`
	UpdatedAt *time.Time  `orm:"auto_now;"`
	DeletedAt *time.Time  `orm:"null;"`
}

type OsqueryTask struct {
	Id   int64
	Name string
	Type string

	Status      int
	ComplatedAt *time.Time `orm:"null"`

	CreatedAt *time.Time `orm:"auto_now_add;"`
	UpdatedAt *time.Time `orm:"auto_now;"`
	DeletedAt *time.Time `orm:"null;"`
}

type OsqueryJob struct {
	Id      int64
	TaskId  int64  `orm:"column(tid);"`
	UUID    string `orm:"column(uuid);"`
	Type    string
	Content string `orm:"type(text);"`

	Status int
	Result string `orm:"type(text);"`

	ComplatedAt *time.Time `orm:"null"`
	CreatedAt   *time.Time `orm:"auto_now_add;"`
	UpdatedAt   *time.Time `orm:"auto_now;"`
	DeletedAt   *time.Time `orm:"null;"`
}

func (job *OsqueryJob) ResultRows() []map[string]string {
	var rows []map[string]string
	if err := json.Unmarshal([]byte(job.Result), &rows); err != nil {
		fmt.Println(err)
	}

	return rows
}

func init() {
	orm.RegisterModel(new(OsqueryAgent), new(OsqueryTaskType), new(OsqueryTask), new(OsqueryJob))
}
