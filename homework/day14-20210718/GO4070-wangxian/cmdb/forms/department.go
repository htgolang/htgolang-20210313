package forms

import (
	"cmdb/models"
	"cmdb/utils"

	"github.com/beego/beego/v2/core/validation"
)

type DepartmentForm struct {
	Id             int64  `form:"id"`
	DepartmentName string `form:"name"`
	Department     *models.Department
}

func (f *DepartmentForm) Valid(v *validation.Validation) {
	v.Required(f.DepartmentName, "name.name").Message("部门名称不能为空")
	v.MaxSize(f.DepartmentName, 128, "naem.name").Message("部门名称不能·超过128个字符")

	if utils.CheckDepartmentExists(f.Id, f.DepartmentName) {
		v.SetError("name", "部门已存在")
	}

	if !v.HasErrors() {
		f.Department = new(models.Department)
		f.Department.Name = f.DepartmentName
		f.Department.Id = f.Id
	}
}

func (f *DepartmentForm) Model2Form(department *models.Department) {
	f.Id = department.Id
	f.DepartmentName = department.Name
}
