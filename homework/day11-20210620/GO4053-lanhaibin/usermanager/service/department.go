package service

import (
	"database/sql"
	"errors"
	"fmt"
	"usermanager/db"
	"usermanager/model"
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
	department := model.Department{}
	rows, err := db.Db.Query(sqlQueryDepartmentById, id)
	if err == nil && rows.Next() {
		err = rows.Scan(&department.Id, &department.Name)
		return department, nil
	}
	return department, errors.New("部门不存在!")
}

func (d DepartmentService) Create(name string) error {
	sql := "insert into department (name) values(?);"
	_, err := d.GetDepartmentByName(name)
	if err == nil {
		return errors.New("部门已存在")
	}
	_, err = db.Db.Exec(sql, name)
	return err
}

func (d DepartmentService) Delete(id int) error {
	_, err := db.Db.Exec(sqlDeleteDepartmentById, id)
	return err
}

func (d DepartmentService) Query(s string) (qs []model.Department) {
	var rows *sql.Rows
	var err error
	if len(s) != 0 {
		s = fmt.Sprintf("%%%s%%", s)
		rows, err = db.Db.Query(sqlQueryDepartmentByName, s)
	} else {
		sql := sqlQueryDepartment
		rows, err = db.Db.Query(sql)
	}
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		department := model.Department{}
		err := rows.Scan(&department.Id, &department.Name)
		if err != nil {
			continue
		}
		qs = append(qs, department)
	}
	return
}

func (d DepartmentService) Modify(id int, name string) error {
	department, err := d.GetDepartmentById(id)
	if err != nil {
		return fmt.Errorf("department id %d not exists", id)
	}
	if len(name) > 0 {
		t, err := d.GetDepartmentByName(name)
		if err == nil && t.Id != id {
			return errors.New("部门已存在")
		}
		department.Name = name
	}
	_, err = db.Db.Exec(sqlUpdateDepartmentById, name, id)
	return err
}

func (d DepartmentService) GetDepartmentByName(name string) (model.Department, error) {
	department := model.Department{}
	rows, err := db.Db.Query(sqlQueryDepartmentByName, name)
	if err == nil && rows.Next() {
		err := rows.Scan(&department.Id, &department.Name)
		return department, err
	}
	return department, errors.New("部门不存在")
}

func (d DepartmentService) GetDepartmentById(id int) (model.Department, error) {
	department := model.Department{}
	rows, err := db.Db.Query(sqlQueryDepartmentById, id)
	if err != nil {
		return department, err
	}
	if rows.Next() {
		err = rows.Scan(&department.Id, &department.Name)
	}
	return department, err
}

var Ds DepartmentService
