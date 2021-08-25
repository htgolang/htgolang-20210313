package main

import (
	"cmdb/models"
	_ "cmdb/routers"
	"cmdb/services"
	"flag"
	"fmt"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var (
	runweb           bool
	syncdb           bool
	force            bool
	verbose          bool
	initUser         bool
	help             bool
	initUsername     string
	initUserPassword string
)

func initDbconn() {
	dbuser, _ := web.AppConfig.String("db::DbUser")
	dbpassword, _ := web.AppConfig.String("db::DbPassword")
	dbhost, _ := web.AppConfig.String("db::DbHost")
	dbport, _ := web.AppConfig.String("db::DbPort")
	dbname, _ := web.AppConfig.String("db::DbName")

	//初始化数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", dbuser, dbpassword, dbhost, dbport, dbname)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	initDbconn()

	flag.BoolVar(&runweb, "web", false, "Start web server")
	flag.BoolVar(&syncdb, "syncdb", false, "Use orm to synchronize goalng structs to the database")
	flag.BoolVar(&force, "force", false, "Whether to delete and recreate a table when it exists in the database")
	flag.BoolVar(&verbose, "verbose", false, "Print the executed sql")
	flag.BoolVar(&initUser, "init-user", false, "Create an initial administrator user")
	flag.BoolVar(&help, "h", false, "Print help message")

	flag.Parse()

	flag.Usage = func() {
		fmt.Println("Usage: cmdb [--web|--syncdb --force --verbose|--init-user]")
		flag.PrintDefaults()
	}

	if help {
		flag.Usage()
		return
	} else if runweb {
		web.Run()
	} else if syncdb {
		orm.RunSyncdb("default", force, verbose)
	} else if initUser {
		fmt.Print("请输入用户名: ")
		fmt.Scan(&initUsername)
		fmt.Print("请输入密码: ")
		fmt.Scan(&initUserPassword)

		user := models.User{Name: initUsername, RoleId: 1}
		user.SetPassword(initUserPassword)
		if err := services.UserService.AddUser(&user); err != nil {
			fmt.Println("初始化管理员失败")
		} else {
			fmt.Println("初始化管理员成功")
		}

	}
}
