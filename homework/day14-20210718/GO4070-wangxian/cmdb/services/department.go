package services

import (
	"cmdb/models"
	"log"

	"github.com/beego/beego/v2/adapter/orm"
)

type departmentService struct{}

var DepartmentService departmentService

func (s *departmentService) QueryDepartment(keyword string) []models.Department {
	var dps []models.Department
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.Department))

	if keyword != "" {
		queryset = queryset.Filter("name__contains", keyword)
	}

	if _, err := queryset.All(&dps); err != nil {
		log.Println("query department error.", err)
		return nil
	}
	return dps

}

func (s *departmentService) QueryDepartmentByID(id int64) *models.Department {
	var dp = models.Department{Id: id}
	ormer := orm.NewOrm()
	if err := ormer.Read(&dp); err != nil {
		log.Println("query department by id error.", err)
		return nil
	}
	return &dp
}

func (s *departmentService) QueryDepartmentByName(name string) *models.Department {
	var dp = models.Department{Name: name}
	ormer := orm.NewOrm()

	if err := ormer.Read(&dp, "name"); err != nil {
		log.Println("query department by name error.", err)
		return nil
	}
	return &dp
}

func (s *departmentService) DeleteDepartment(id int64) {
	dp := models.Department{Id: id}
	ormer := orm.NewOrm()

	if _, err := ormer.Delete(&dp); err != nil {
		log.Println("delete department error.", err)
	}
}

func (s *departmentService) AddDepartment(department *models.Department) error {
	ormer := orm.NewOrm()
	if _, err := ormer.Insert(department); err != nil {
		log.Println("add department error.", err)
		return err
	}
	return nil
}

func (s *departmentService) EditDepartment(department *models.Department) error {
	ormer := orm.NewOrm()
	if _, err := ormer.Update(department); err != nil {
		log.Println("update department error.", err)
		return err
	}
	return nil
}
