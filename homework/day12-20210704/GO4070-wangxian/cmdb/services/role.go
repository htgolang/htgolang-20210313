package services

import (
	"cmdb/db"
	"cmdb/models"
	"log"
)

func QueryRole() []models.Role {
	var roles []models.Role
	sql := `select id,name,modules from role;`
	rows, err := db.DbConn.Query(sql)
	if err != nil {
		log.Printf("query_role: query data from db error. %s", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.Id, &role.Name, &role.Modules); err != nil {
			log.Printf("quer_role: scan data from rows error. %s", err)
			continue
		}
		roles = append(roles, role)
	}
	return roles
}

func QueryRoleByID(id int64) *models.Role {
	// fmt.Println(id)
	var role models.Role
	sql := `select id,name,modules from role where id=?;`
	err := db.DbConn.QueryRow(sql, id).Scan(&role.Id, &role.Name, &role.Modules)
	if err != nil {
		log.Printf("query_role_byid: scan data from row error. %s", err)
		return nil
	}
	return &role
}
