package services

import (
	"cmdb/db"
	"cmdb/models"
	"database/sql"
	"fmt"
	"log"
)

func QueryDepartment(keyword string) []models.Department {
	var dps []models.Department
	var rows *sql.Rows
	var err error
	if keyword == "" {
		sql := `select id,name from department;`
		rows, err = db.DbConn.Query(sql)
	} else {
		sql := `select id,name from department where name like ?;`
		rows, err = db.DbConn.Query(sql, fmt.Sprintf("%%%s%%", keyword))
	}

	if err != nil {
		log.Printf("query department from db error. %s", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var dp models.Department
		err := rows.Scan(&dp.Id, &dp.Name)
		if err != nil {
			log.Printf("scan department from rows error. %s", err)
			continue
		}
		dps = append(dps, dp)
	}
	return dps

}

func QueryDepartmentByID(id int64) *models.Department {
	var dp models.Department
	sql := `select id,name from department where id=?;`
	if err := db.DbConn.QueryRow(sql, id).Scan(&dp.Id, &dp.Name); err == nil {
		return &dp
	} else {
		log.Printf("query department by id:scan data from rwo error. %s", err)
		return nil
	}
}

func DeleteDepartment(id int64) {
	sql := `delete from department where id=?;`
	_, err := db.DbConn.Exec(sql, id)
	if err != nil {
		log.Printf("delete department: delete data from db error. %s", err)
	}
}

func AddDepartment(name string) error {
	sql := `insert into department(name,create_at,update_at) values(?,now(),now());`
	_, err := db.DbConn.Exec(sql, name)
	return err
}

func EditDepartment(id int64, name string) error {
	sql := `update department set name=? where id=?;`
	_, err := db.DbConn.Exec(sql, name, id)
	if err != nil {
		log.Printf("edit department: update data to db error. %s", err)
	}
	return err
}
