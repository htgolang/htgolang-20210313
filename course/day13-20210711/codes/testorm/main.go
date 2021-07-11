package main

// 导入包
// orm
// mysql 驱动
import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // dataset/sql
)

// 用beego orm 显示设置主键
// ID属性需要设置 int64 或者显示指定pk
/**
标签名: orm 标签值用;分隔
指定主键: pk
自动增长: auto
列名: column(name)
字符串类型: 默认varchar(255)
				长度修改 size(length)
			指定其他类型: type(text)
是否允许为null: 默认不允许, null
唯一: unquie
索引: index

时间:
	type(date);
	// 不体现再sql中，实在orm执行过程中应用
	auto_now; // 每次更新的时候自动设置属性为当前时间
	auto_now_add; //当数据创建时设置为当前时间
*/
type Account struct {
	ID           int64  `orm:"pk;auto;column(id);"`
	Name         string `orm:"size(64);"`
	Password     string `orm:"size(1024);"`
	Birthday     *time.Time
	Telephone    string
	Email        string
	Addr         string `orm:"size(4096);default(中国);"`
	Status       int8   `orm:"default(1);description(状态);"`
	RoleId       int64
	DepartmentId int64
	CreatedAt    *time.Time `orm:"auto_now_add;"`
	UpdatedAt    *time.Time `orm:"auto_now;"`
	DeletedAt    *time.Time `orm:"null;"`
	Description  string     `orm:"type(text);"`
	Sex          bool
	Height       float32 `orm:"digits(10);decimals(2);"`
	Weight       float64
}

// 指定关联的表
func (account *Account) TableName() string {
	return "act"
}

// 指定索引
func (account *Account) TableIndex() [][]string {
	return [][]string{
		{"name"},
		{"telephone", "email"},
	}
}

func main() {
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

	// 使用
	// 表结构同步（库需要提前创建）
	// databaseName: 同步数据库
	// force: 如果表存在则删除
	// verbose: 打印执行的sql语句
	orm.RunSyncdb(databaseName, true, true)
}
