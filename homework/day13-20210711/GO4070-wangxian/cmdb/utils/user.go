package utils

import (
	"cmdb/models"
	"fmt"

	"github.com/beego/beego/v2/adapter/orm"
)

func CheckUserExists(uid int64, field, value string) bool {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.User)).Filter(fmt.Sprintf("%s__exact", field), value)

	if uid >= 0 {
		queryset = queryset.Exclude("id__exact", uid)
	}

	return queryset.Exist()
}
