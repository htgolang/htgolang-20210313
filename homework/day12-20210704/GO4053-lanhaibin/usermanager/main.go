package main

import (
	_ "usermanager/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}
