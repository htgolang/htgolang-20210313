package services

import (
	"cmdb/models"
	"log"

	"github.com/beego/beego/v2/client/orm"
)

// 对外都调用UserService
var UserService = new(userService)

type userService struct {
}

func (s *userService) ValidateLogin(name string, password string) *models.User {
	user := s.GetByName(name)
	if user != nil && user.ValidPassword(password) {
		return user
	}
	return nil
}

func (s *userService) Create(user *models.User) error {
	ormer := orm.NewOrm()
	_, err := ormer.Insert(user)
	return err
}

func (s *userService) GetById(id int64) *models.User {
	user := &models.User{Id: id}
	ormer := orm.NewOrm()
	if err := ormer.Read(user); err == nil {
		return user
	}
	return nil
}

func (s *userService) GetByName(name string) *models.User {
	user := &models.User{Name: name}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

func (s *userService) Query(q string) []models.User {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.User))
	if q != "" {
		cond := orm.NewCondition()
		cond = cond.Or("name__icontains", q)
		cond = cond.Or("telephone__icontains", q)
		cond = cond.Or("email__icontains", q)
		cond = cond.Or("addr__icontains", q)
		cond = cond.Or("description__icontains", q)
		queryset = queryset.SetCond(cond)
	}
	users := make([]models.User, 0)
	if _, err := queryset.All(&users); err != nil {
		log.Println(err)
	}
	return users
}

func (s *userService) Delete(id int64) {
	ormer := orm.NewOrm()
	if _, err := ormer.Delete(&models.User{Id: id}); err != nil {
		log.Println(err)
	}
}

func (s *userService) ModifyPassword(user *models.User) {
	ormer := orm.NewOrm()
	ormer.Update(user, "Password") // 只更新密码字段
}
