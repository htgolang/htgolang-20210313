package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	_ "webusermanage/router"
	"webusermanage/config"

)

func main() {

	
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/login/", http.StatusFound)
	// })
	http.ListenAndServe(config.Addr, nil)

}
