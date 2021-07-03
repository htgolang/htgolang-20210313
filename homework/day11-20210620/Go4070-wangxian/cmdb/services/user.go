package services

import (
	"cmdb/db"
	"cmdb/models"
	"cmdb/utils"
	"database/sql"
	"fmt"
	"log"
)

func QueryUser(keyWord string) []models.User {
	var users []models.User
	var rows *sql.Rows
	var err error
	if keyWord == "" {
		sql := `select id,name,birthday,telephone,email,addr,description,sex from user;`
		rows, err = db.DbConn.Query(sql)

	} else {
		sql := `select id,name,birthday,telephone,email,addr,description,sex from user where name like ?;`
		rows, err = db.DbConn.Query(sql, fmt.Sprintf("%%%s%%", keyWord))
	}

	if err != nil {
		log.Printf("query users from db error.%s", err)
		return nil
	}

	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Name, &user.Birthday, &user.Telephone, &user.Email, &user.Addr, &user.Description, &user.Sex)
		if err != nil {
			log.Printf("query_user: scan data from rows error.%s", err)
			continue
		}
		users = append(users, user)
	}

	return users
}

func DeleteUser(id int64) {
	sql := `delete from user where id=?;`
	db.DbConn.Exec(sql, id)
}

func AddUser(username, sex, birthday, tel, email, addr, description, password, confirmpw string, roleid int64) error {
	var u models.User
	if password != confirmpw {
		return fmt.Errorf("两次密码不一致")
	}

	if utils.CheckUserExists(-1, "name", username) {
		return fmt.Errorf("用户名已注册")
	}

	if utils.CheckUserExists(-1, "telephone", tel) {
		return fmt.Errorf("手机号已注册")
	}

	if utils.CheckUserExists(-1, "email", email) {
		return fmt.Errorf("邮箱已注册")
	}
	u.SetPassword(password)

	sql := `insert into user(name, password, sex, birthday, telephone, email, addr, description, role_id, create_at, update_at) values(?, ?, ?, ?, ?, ?, ?, ?, ?, now(), now());`
	_, err := db.DbConn.Exec(sql, username, u.Password, sex, birthday, tel, email, addr, description, roleid)
	if err != nil {
		log.Printf("add_user：insert user to db error.%s", err)
		return fmt.Errorf("添加用户失败")
	}

	return nil
}

func QueryUserByID(id int64) *models.User {
	var user models.User
	sql := `select id,name,birthday,telephone,email,addr,description,sex,role_id from user where id=?;`
	err := db.DbConn.QueryRow(sql, id).Scan(&user.Id, &user.Name, &user.Birthday, &user.Telephone, &user.Email, &user.Addr, &user.Description, &user.Sex, &user.RoleId)
	if err != nil {
		log.Printf("query_user_by_id: scan data from row error.%s", err)
		return nil
	}
	return &user
}

func EditUser(uid int64, username, sex, birthday, tel, email, addr, description string, roleid int64) error {

	if utils.CheckUserExists(uid, "name", username) {
		return fmt.Errorf("用户名已注册")
	}

	if utils.CheckUserExists(uid, "telephone", tel) {
		return fmt.Errorf("手机号已注册")
	}

	if utils.CheckUserExists(uid, "email", email) {
		return fmt.Errorf("邮箱已注册")
	}

	sql := `update user set name=?,sex=?,birthday=?,telephone=?,email=?,addr=?,description=?,role_id=?, update_at=now() where id=?;`
	_, err := db.DbConn.Exec(sql, username, sex, birthday, tel, email, addr, description, roleid, uid)
	if err != nil {
		log.Printf("edit_user: update user to db error. %s", err)
		return fmt.Errorf("更新用户信息失败")
	}

	return nil
}

func ChangePw(uid int64, oldPassword, newPassword, confirmpw string) error {
	var user models.User
	sql := `select id,name,password from user where id=?;`
	err := db.DbConn.QueryRow(sql, uid).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		log.Printf("changepw: scan data from row error. %s", err)
		return fmt.Errorf("密码修改失败，请重试")
	}

	if !user.ValidPassword(oldPassword) {
		return fmt.Errorf("原密码错误")
	}

	if newPassword != confirmpw {
		return fmt.Errorf("两次密码不一致")
	}

	if user.ValidPassword(newPassword) {
		return fmt.Errorf("新密码不能和原密码一致")
	}

	user.SetPassword(newPassword)
	sql = `update user set password=? where id=? and name=?;`
	_, err = db.DbConn.Exec(sql, user.Password, user.Id, user.Name)
	if err != nil {
		log.Printf("changepw: update user data to db error. %s", err)
		return fmt.Errorf("密码修改失败，请重试")
	}
	return nil
}
