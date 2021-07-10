package controllers

import (
	"net/http"
	"usermanager/service"
)

type DepartmentCreateController struct {
	AuthController
}

func (this *DepartmentCreateController) Get() {
	this.Data["Err"] = ""
	this.TplName = "department/department_create.html"
}

func (this *DepartmentCreateController) Post() {
	name := this.GetString("name")
	err := service.Ds.Create(name)
	if err != nil {
		this.Data["Err"] = err.Error()
		this.TplName = "department/department_create.html"
		return
	}
	this.Redirect("/department", http.StatusFound)
}

type DepartmentDeleteController struct {
	AuthController
}

func (this *DepartmentDeleteController) Get() {
	id, _ := this.GetInt("id")
	service.Ds.Delete(id)
	this.Redirect("/department", http.StatusFound)
}

type DepartmentQueryController struct {
	AuthController
}

func (this *DepartmentQueryController) Get() {
	s := this.GetString("q")
	departmentList := service.Ds.Query(s)
	this.Data["Departments"] = departmentList
	this.TplName = "department/department.html"
}

type DepartmentModifyController struct {
	AuthController
}

func (this *DepartmentModifyController) Get() {
	id, _ := this.GetInt("id")
	department, err := service.Ds.Get(id)
	if err != nil {
		this.Redirect("/department", http.StatusFound)
		return
	}
	this.Data["Department"] = department
	this.TplName = "department/department_modify.html"
}

func (this *DepartmentModifyController) Post() {
	id, _ := this.GetInt("id")
	name := this.GetString("name")

	service.Ds.Modify(id, name)
	this.Redirect("/department", http.StatusFound)
}
