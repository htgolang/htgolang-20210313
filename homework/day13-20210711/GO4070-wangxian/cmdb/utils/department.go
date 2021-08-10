package utils

import (
	"cmdb/models"

	"github.com/beego/beego/v2/client/orm"
)

func CheckDepartmentExists(dpid int64, name string) bool {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.Department)).Filter("name__exact", name)
	if dpid > 0 {
		queryset = queryset.Exclude("id__exact", dpid)
	}

	return queryset.Exist()
}
