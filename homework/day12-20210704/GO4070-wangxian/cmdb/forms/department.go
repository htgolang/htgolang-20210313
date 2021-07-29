package forms

import (
	"cmdb/utils"

	"github.com/beego/beego/v2/core/validation"
)

type DepartmentForm struct {
	Id             int64  `form:"id"`
	DepartmentName string `form:"name"`
}

func (f *DepartmentForm) Valid(v *validation.Validation) {
	v.Required(f.DepartmentName, "name.name").Message("部门名称不能为空")

	if utils.CheckDepartmentExists(f.Id, f.DepartmentName) {
		v.SetError("name", "部门已存在")
	}
}
