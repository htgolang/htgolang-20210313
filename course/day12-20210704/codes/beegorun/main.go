package main

import (
	_ "beegorun/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run() // http.ListenAndServer()
}
