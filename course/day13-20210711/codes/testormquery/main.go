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
	ormer := orm.NewOrm()
	// select id,name from table where xxx order by limit offset group by
	queryset := ormer.QueryTable(new(Account))

	// queryset.Count() // select count(*)
	// queryset.All()   // select columns from

	// where key condition value;
	// Filter("key__condition", value)
	// id = 1 Filter("id__exact", 1) Filter("id", 1)
	// name = "kk" Filter("name__exact", "kk") Filter("name", "kk")
	/* codition:
		(i)exact => (不区分大小写的) =
	关系运算
		gt => >
		lt => <
		gte => >=
		lte => <=
	字符串:
		以开头 (i)startswith
		以结尾 (i)endswith
		包含 (i)contains
	in
	isnull

	and

	Filter(id).Filter(name) 过滤出
	Exclude("id", 1).Exclude() 排除掉

	*/

	queryset.Filter("id", 1).
		Filter("name__icontains", "kk").
		Filter("addr__startswith", "北京").
		Filter("deleted_at__isnull", true).
		Filter("status__in", []int{1, 2}).
		Exclude("birthday__lte", "2020-01-02").
		Count()

	// (a and not b) and ((c or d) or not (e and f))
	// Condition
	// And
	// Or
	// AndCond
	// OrCond
	// AndNot
	// OrNot
	// AndNotCond
	// OrNotCond

	// (id = 1 and not name like %xxxx%) and ( ( addr like %yyyy% or status in [1, 2, 3]) or not (birthday >= 2000-01-01 and description like '%zzz'))

	//  ( T0.`id` = ? AND NOT T0.`name` LIKE BINARY ? ) AND ( ( T0.`addr` LIKE BINARY ? OR T0.`status` IN (?, ?, ?) ) OR NOT ( T0.`birthday` >= ? AND T0.`description` LIKE BINARY ? ) ) ] - `1`, `%xxxx%`, `%yyyy%`, `1`, `2`, `3`, `2000-01-01 00:00:00`, `%zzz`
	cond1 := orm.NewCondition()
	cond1 = cond1.And("id", 1)
	cond1 = cond1.AndNot("name__contains", "xxxx")

	cond2 := orm.NewCondition()
	cond2 = cond2.And("addr__contains", "yyyy").Or("status__in", []int{1, 2, 3})

	cond3 := orm.NewCondition()
	cond3 = cond3.And("birthday__gte", "2000-01-01").And("description__endswith", "zzz")

	cond4 := orm.NewCondition()
	cond4 = cond4.AndCond(cond2).OrNotCond(cond3)

	cond5 := orm.NewCondition()
	cond5 = cond5.AndCond(cond1).AndCond(cond4)

	queryset.SetCond(cond5).Count()

	// order by

	// 指定多个列 column 升序 -column 降序
	accounts := make([]Account, 0)
	queryset.OrderBy("id", "-name").All(&accounts)
	// limit offset
	queryset.OrderBy("id").Limit(3).Offset(10).All(&accounts)

	var account Account
	queryset.OrderBy("id").One(&account)

	var params []orm.Params
	queryset.OrderBy("id").Values(&params)
	for _, param := range params {
		fmt.Println(param)
	}

	var paramsList []orm.ParamsList

	queryset.OrderBy("id").ValuesList(&paramsList, "Name", "Id")
	for _, param := range paramsList {
		fmt.Println(param)
	}

	// 批量更新
	// Update()
	// 批量删除
	// Delete()
	fmt.Println(queryset.Filter("id__lte", 60).Delete())
	fmt.Println(queryset.Filter("id__lte", 62).Update(
		orm.Params{
			"addr":        "西安",
			"description": "测试",
			"status":      orm.ColValue(orm.ColAdd, 1),
		},
	))
}
