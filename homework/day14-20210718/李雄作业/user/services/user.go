package services

import (
	"log"
	"user/forms"
	"user/models"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

// 查询数据库获取用户
func GetUsers() []*models.User {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&models.User{})
	var users []*models.User
	queryset.All(&users)
	return users
}

func AddUser(name string, password string, addr string, sex bool) {
	// 密码加密
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	user := &models.User{
		Name:     name,
		Password: string(hashed),
		Addr:     addr,
		Sex:      sex,
	}
	ormer := orm.NewOrm()
	ormer.Insert(user)
}

// 修改密码函数
func ChangeUser(name string, password string, addr string, sex bool) bool {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	NewUser := &models.User{
		Name:     name,
		Password: string(hashed),
		Addr:     addr,
		Sex:      sex,
	}
	ormer := orm.NewOrm()
	user := &models.User{Name: name}
	if err := ormer.Read(user); err == nil {
		log.Fatal("密码修改错误")
		return false
	} else {
		_, err = ormer.Update(NewUser)
		if err != nil {
			log.Fatal("密码修改错误")
			return false
		}
	}
	return true
}

func DeleteUser(id int64) {
	ormer := orm.NewOrm()
	ormer.Delete(&models.User{ID: id})
}

func GetUserById(pk int64) *models.User {
	ormer := orm.NewOrm()
	user := &models.User{ID: pk}
	if err := ormer.Read(user); err == nil {
		return user
	}
	return nil
}

// 登录时获取用户信息 对接数据库
func GetUserByName(name string) *models.User {
	user := &models.User{
		Name: name,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

// 对比用户输入的用户信息和 数据库存储的数据信息 来判断用户是否存在/密码是否正确
func Auth(form *forms.LoginFrom) *models.User {
	// 通过用户名去查询用户信息，至少包含密码hash值
	// 将从数据库查询到的数据进行处理
	if user := GetUserByName(form.Username); user == nil {
		// 查询不到用户，证明用户不存在，或者输入有误 直接返回nil
		return nil
	} else {
		// 查询到数据了 密码检查
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err == nil {
			// 成功 返回用户信息
			return user
		} else {
			// 用户名或密码错误  直接返回nil
			return nil
		}
	}
}
