package main

import (
	"log"

	_ "user/routers"
	"user/services"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := beego.AppConfig.String("mysql::DSN") // 通过mysql的配置读取
	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		log.Fatal(err) // 记录错误并退出
		// default 数据库链接别名  多个数据建库链接
	}
	orm.RunSyncdb("default", false, true)
	if user := services.GetUserByName("admin"); user == nil {
		log.Print("创建管理员账号")
		services.AddUser("admin", "123@abc", "深圳", true)
	}

	addr := ":8888"
	// beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "192.168.56.103:6379"

	beego.Run(addr)
}
