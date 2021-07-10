package main

import (
	_ "webusermanage/routers"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/login/", http.StatusFound)
	// })
	// http.ListenAndServe(config.Addr, nil)
	web.Run()

}
