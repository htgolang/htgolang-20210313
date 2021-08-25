package services

import (
	"cmdb/models"
	"log"

	"github.com/beego/beego/v2/client/orm"
)

type userService struct{}

var UserService = new(userService)

func (s *userService) QueryUser(keyWord string) []models.User {
	var users []models.User
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.User))
	if keyWord != "" {
		cond := orm.NewCondition()
		cond = cond.And("name__contains", keyWord)
		cond = cond.Or("telephone__contains", keyWord)
		cond = cond.Or("email__contains", keyWord)
		cond = cond.Or("addr__contains", keyWord)
		cond = cond.Or("description__contains", keyWord)
		queryset = queryset.SetCond(cond)
	}

	if _, err := queryset.All(&users); err != nil {
		log.Println("query users error by orm error.", err)
	}

	return users
}

func (s *userService) DeleteUser(id int64) {
	var user = models.User{Id: id}
	ormer := orm.NewOrm()
	if _, err := ormer.Delete(&user); err != nil {
		log.Println("delete user error.", err)
	}
}

func (s *userService) AddUser(user *models.User) error {
	ormer := orm.NewOrm()
	if _, err := ormer.Insert(user); err != nil {
		log.Println("add user error by orm error.", err)
		return err
	}
	return nil
}

func (s *userService) QueryUserByID(id int64) *models.User {
	var user = models.User{Id: id}
	ormer := orm.NewOrm()
	if err := ormer.Read(&user); err != nil {
		log.Println("query user_by_id by orm error.", err)
		return nil
	}
	return &user

}

func (s *userService) QueryUserByName(name string) *models.User {
	var user = models.User{Name: name}
	ormer := orm.NewOrm()
	if err := ormer.Read(&user, "name"); err != nil {
		log.Println("query user by name error.", err)
		return nil
	}
	return &user
}

func (s *userService) EditUser(user *models.User) error {
	ormer := orm.NewOrm()
	if _, err := ormer.Update(user, "name", "sex", "birthday", "telephone", "email", "addr", "role_id", "description"); err != nil {
		log.Println("modify user error.", err)
		return err
	}
	return nil
}

func (s *userService) ModifyPw(user *models.User) error {
	ormer := orm.NewOrm()
	if _, err := ormer.Update(user, "password"); err != nil {
		log.Println("update user password error.", err)
		return err
	}
	return nil
}
