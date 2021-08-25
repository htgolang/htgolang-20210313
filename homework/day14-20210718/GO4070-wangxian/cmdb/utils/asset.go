package utils

import (
	"cmdb/models"

	"github.com/beego/beego/v2/client/orm"
)

func CheckAssetExists(id int64, ip string) bool {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.Asset)).Filter("ip__exact", ip)

	if id > 0 {
		queryset = queryset.Exclude("id__exact", id)
	}

	return queryset.Exist()

}
