package forms

import (
	"cmdb/models"
	"strings"

	"github.com/beego/beego/v2/core/validation"
)

type OsqueryTaskForm struct {
	Name  string   `form:"name"`
	UUIDs []string `form:"uuids"`
	Type  string   `form:"type"`
}

func (f *OsqueryTaskForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)

	v.MinSize(f.Name, 3, "name.name.name").Message("名称长度必须为3~32")
	v.MaxSize(f.Name, 32, "name.name.name").Message("名称长度必须为3~32")
	v.MinSize(f.UUIDs, 1, "targets.targets.targets").Message("扫描目标不能为空")
}

func (f *OsqueryTaskForm) ToModel() *models.OsqueryTask {
	return &models.OsqueryTask{
		Name: f.Name,
		Type: f.Type,
	}
}
