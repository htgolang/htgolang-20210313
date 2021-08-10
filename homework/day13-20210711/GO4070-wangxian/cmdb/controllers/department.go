package controllers

import (
	"cmdb/forms"
	"cmdb/models"
	"cmdb/services"
	"log"
	"net/http"

	"github.com/beego/beego/v2/adapter/validation"
)

type DepartmentController struct {
	AuthenticationController
}

func (c *DepartmentController) Query() {
	keyword := c.GetString("q")
	departments := services.DepartmentService.QueryDepartment(keyword)

	c.Data["Departments"] = departments
	c.TplName = "department/list.html"
}

func (c *DepartmentController) Delete() {
	id, _ := c.GetInt64("id")
	services.DepartmentService.DeleteDepartment(id)
	c.Redirect("/department/query/", http.StatusFound)
}

func (c *DepartmentController) Add() {
	var form forms.DepartmentForm
	var valid validation.Validation

	if c.Ctx.Input.IsPost() {
		err := c.ParseForm(&form)
		if err == nil {
			if b, err := valid.Valid(&form); err != nil {
				log.Println("valid add department data error.", err)
				valid.SetError("department", "验证数据错误")
			} else if b {
				err := services.DepartmentService.AddDepartment(form.Department)
				if err == nil {
					c.Redirect("/department/query/", http.StatusFound)
					return
				}
				log.Println("add department error.", err)
				valid.SetError("department", "添加部门失败")
			}
		} else {
			log.Println("parse add department data error.", err)
			valid.SetError("department", "提交数据错误")
		}
	}

	c.Data["ErrMsgs"] = valid.ErrorMap()
	c.TplName = "department/add.html"
}

func (c *DepartmentController) Edit() {
	var form forms.DepartmentForm
	var valid validation.Validation
	var department *models.Department = &models.Department{}

	if c.Ctx.Input.IsGet() {
		id, _ := c.GetInt64("id")
		department = services.DepartmentService.QueryDepartmentByID(id)
	}

	if c.Ctx.Input.IsPost() {
		err := c.ParseForm(&form)
		if err == nil {
			if b, err := valid.Valid(&form); err != nil {
				log.Println("valid edit department data error.", err)
				valid.SetError("department", "验证数据错误")
			} else if b {
				err := services.DepartmentService.EditDepartment(form.Department)
				if err == nil {
					c.Redirect("/department/query/", http.StatusFound)
					return
				}
				log.Println("edit department error.", err)
				valid.SetError("department", "添加部门失败")
			} else {
				department.Id = form.Id
				department.Name = form.DepartmentName
			}
		} else {
			log.Println("parse edit department data error.", err)
			valid.SetError("department", "提交数据错误")
		}
	}

	c.Data["Department"] = department
	c.Data["ErrMsgs"] = valid.ErrorMap()
	c.TplName = "department/edit.html"
}
