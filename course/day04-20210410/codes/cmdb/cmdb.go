package main

import (
	"cmdb/user" // 同module => modulename/ 目录路径
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	fmt.Println("hello world")
	user.Add()
	web.Run()
}
