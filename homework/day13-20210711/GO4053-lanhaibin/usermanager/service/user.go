package service

import (
	"errors"
	"fmt"
	"time"
	"usermanager/model"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

const (
	sqlQueryUserById      = "select id, name, brithday, sex, addr, tel, password from user where id=?;"
	sqlQueryUserByName    = "select id, name, brithday, sex, addr, tel, password from user where name like ?;"
	sqlQueryUser          = "select id, name, brithday, sex, addr, tel, password from user;"
	sqlDeleteUserById     = "delete from user where id=?;"
	sqlUpdateUserById     = "update user set name=?, brithday=?, sex=?,addr=?,tel=? where id=?;"
	sqlUpdateUserPassword = "update user set password=? where id=?;"
)

type UserService struct{}

func (u UserService) Get(id int) (model.User, error) {
	o := orm.NewOrm()
	user := model.User{Id: id}
	err := o.Read(&user)
	return user, err
}

func (u UserService) Create(name string, brithday *time.Time, sex bool, addr, tel, password string) error {
	if len(name) == 0 || len(password) == 0 {
		return errors.New("用户名和密码不能为空")
	}
	user, err := u.GetUserByName(name)
	if err == nil {
		return errors.New("用户已存在")
	}
	user = model.User{
		Name:     name,
		Brithday: brithday,
		Sex:      sex,
		Addr:     addr,
		Tel:      tel,
	}
	user.SetPassword(password)
	// _, err = db.Db.Exec(sql, name, brithday, sex, addr, tel, user.Password)
	o := orm.NewOrm()
	_, err = o.Insert(&user)
	return err
}

func (u UserService) Delete(id int) error {
	user := model.User{Id: id}
	o := orm.NewOrm()
	count, err := o.Delete(&user)
	if count != 1 || err != nil {
		return errors.New("删除失败!")
	}
	return nil
}

func (u UserService) Query(s string) []model.User {
	var users []model.User
	o := orm.NewOrm()
	qs := o.QueryTable(new(model.User))
	qs.Filter("name__contains", s).All(&users)
	return users
}

func (u UserService) Modify(id int, name string, brithday *time.Time, sex, addr, tel string) error {
	user, err := u.GetUserById(id)
	if err != nil {
		return fmt.Errorf("user id %d not exists", id)
	}
	user.Name = name
	user.Brithday = brithday
	user.Sex = sex == "1"
	user.Addr = addr
	user.Tel = tel

	o := orm.NewOrm()
	_, err = o.Update(&user)
	return err
}

func (u UserService) GetUserByName(name string) (model.User, error) {
	user := model.User{Name: name}
	o := orm.NewOrm()
	err := o.Read(&user, "Name")
	return user, err

}

func (u UserService) GetUserById(id int) (model.User, error) {
	return u.Get(id)
}

func (u UserService) ChangePassword(id int, password string) error {
	user, _ := u.GetUserById(id)
	user.SetPassword(password)
	o := orm.NewOrm()
	_, err := o.Update(&user)
	if err != nil {
		logs.Error(err.Error())
		return errors.New("修改密码失败")
	}
	return nil
}

var Us UserService
