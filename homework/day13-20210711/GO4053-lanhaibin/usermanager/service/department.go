package service

import (
	"errors"
	"fmt"
	"usermanager/model"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

const (
	sqlQueryDepartmentById   = "select id, name from department where id=?;"
	sqlQueryDepartmentByName = "select id, name from department where name like ?;"
	sqlQueryDepartment       = "select id, name from department;"
	sqlDeleteDepartmentById  = "delete from department where id=?;"
	sqlUpdateDepartmentById  = "update department set name=? where id=?;"
)

type DepartmentService struct{}

func (d DepartmentService) Get(id int) (model.Department, error) {
	department := model.Department{Id: id}
	o := orm.NewOrm()
	err := o.Read(&department)
	if err != nil {
		logs.Error(err.Error())
		return department, errors.New("部门不存在!")
	}
	return department, nil
}

func (d DepartmentService) Create(name string) error {
	department := model.Department{Name: name}
	o := orm.NewOrm()
	_, err := o.Insert(&department)
	if err != nil {
		logs.Error(err.Error())
		return errors.New("新建部门失败")
	}
	return nil
}

func (d DepartmentService) Delete(id int) error {
	department := model.Department{Id: id}
	o := orm.NewOrm()
	count, err := o.Delete(&department)
	if count != 1 || err != nil {
		logs.Error(err.Error())
		return errors.New("删除失败")
	}
	return nil
}

func (d DepartmentService) Query(s string) []model.Department {
	var departments []model.Department
	qs := orm.NewOrm().QueryTable(new(model.Department))
	qs.Filter("name__contains", s).All(&departments)
	return departments
}

func (d DepartmentService) Modify(id int, name string) error {
	department, err := d.GetDepartmentById(id)
	if err != nil {
		return fmt.Errorf("department id %d not exists", id)
	}

	t, err := d.GetDepartmentByName(name)
	if err == nil && t.Id != id {
		return errors.New("部门已存在")
	}
	department.Name = name

	_, err = orm.NewOrm().Update(&department)
	if err != nil {
		logs.Error(err.Error())
		return errors.New("修改失败")
	}
	return nil
}

func (d DepartmentService) GetDepartmentByName(name string) (model.Department, error) {
	department := model.Department{Name: name}
	o := orm.NewOrm()
	err := o.Read(&department)
	if err != nil {
		logs.Error(err.Error())
		return department, errors.New("部门不存在")
	}
	return department, nil
}

func (d DepartmentService) GetDepartmentById(id int) (model.Department, error) {
	return d.Get(id)
}

var Ds DepartmentService
