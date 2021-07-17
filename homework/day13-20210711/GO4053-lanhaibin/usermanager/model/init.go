package model

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbName, _ := web.AppConfig.String("database::name")
	user, _ := web.AppConfig.String("database::user")
	password, _ := web.AppConfig.String("database::password")
	host, _ := web.AppConfig.String("database::host")
	port, _ := web.AppConfig.Int("database::port")

	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", user, password, host, port, dbName))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Department))
}
