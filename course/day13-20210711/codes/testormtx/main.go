package main

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // dataset/sql
)

type Account struct {
	ID           int64 `orm:"column(id);"`
	Name         string
	Password     string
	Birthday     *time.Time
	Telephone    string
	Email        string
	Addr         string
	Status       int8
	RoleId       int64
	DepartmentId int64
	CreatedAt    *time.Time `orm:"auto_now_add;"`
	UpdatedAt    *time.Time `orm:"auto_now";`
	DeletedAt    *time.Time
	Description  string
	Sex          bool
}

func (account *Account) TableName() string {
	return "act"
}

func main() {
	orm.Debug = true
	databaseName := "default"
	driverName := "mysql"
	dsn := "golang:golang@2021@tcp(10.0.0.2:3306)/cmdb?charset=utf8mb4&parseTime=true"

	// 注册驱动到orm
	orm.RegisterDriver(driverName, orm.DRMySQL) // 类型名(自定义), 类型(由orm指定的)

	// 数据库(连接池) 连接(池)名称 使用的驱动名称(orm驱动类型名) 数据库的配置信息
	// beego 指定 default数据库名称
	err := orm.RegisterDataBase(databaseName, driverName, dsn)
	if err != nil {
		log.Fatal(err)
	}

	// 定义结构
	// 注册模型
	orm.RegisterModel(new(Account))
	ormer := orm.NewOrm()

	handle := func() error {
		return nil // fmt.Errorf("error")
	}

	if tx, err := ormer.Begin(); err == nil {
		tx.Raw("delete from act").Exec()
		// err := tx.Insert()
		if err := handle(); err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}

}
