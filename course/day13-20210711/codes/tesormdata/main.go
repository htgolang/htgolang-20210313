package main

import (
	"fmt"
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
	// orm.RunSyncdb(databaseName, true, true)
	ormer := orm.NewOrm()
	now := time.Now()
	// // 增
	// account := Account{
	// 	Name:      "xxxxxxxxxxxxxkk",
	// 	Password:  "123456",
	// 	Birthday:  &now,
	// 	Telephone: "123",
	// 	Email:     "456@test.com",
	// }

	// // 插入数据并将数据库自动生成的值回填
	// id, err := ormer.Insert(&account)
	// fmt.Println(id, err, account)

	// id, err = ormer.Insert(&account)
	// fmt.Println(id, err, account)

	// 删
	// deleteAccount := &Account{ID: 2}
	// num, err := ormer.Delete(deleteAccount)
	// fmt.Println(err, num)

	// deleteAccount = &Account{Name: "abc", Telephone: ""}
	// num, err = ormer.Delete(deleteAccount, "Name", "Telephone")
	// fmt.Println(err, num)
	// 查
	// 查一个结果
	// accountDetail := &Account{ID: 3}
	// err = ormer.Read(accountDetail, "Name", "Telephone")
	// fmt.Println(err, accountDetail)
	// // 查列表
	// // 改
	// accountDetail.Addr = "北京"
	// num, err := ormer.Update(accountDetail)
	// fmt.Println(err, num)

	// // 查询列表
	// // QueryTable
	// queryset := ormer.QueryTable(new(Account))

	// fmt.Println(queryset.Count())

	// accounts := make([]Account, 0)

	// queryset.All(&accounts)
	// fmt.Println(accounts)

	// accounts := make([]Account, 0)
	// for i := 0; i < 10; i++ {
	// 	// 增
	// 	account := Account{
	// 		Name:      fmt.Sprintf("kk-%s", i),
	// 		Password:  "123456",
	// 		Birthday:  &now,
	// 		Telephone: "123",
	// 		Email:     "456@test.com",
	// 	}
	// 	accounts = append(accounts, account)

	// 	// 插入数据并将数据库自动生成的值回填
	// 	// id, err := ormer.Insert(&account)
	// 	// fmt.Println(id, err, account)

	// }
	// // 批次提交的数量 提交数据的列表(切片)
	// fmt.Println(ormer.InsertMulti(8, accounts))

	account := &Account{
		Name:     "kk",
		Birthday: &now,
	}

	created, id, err := ormer.ReadOrCreate(account, "Name")
	fmt.Println(created, id, err)
}
