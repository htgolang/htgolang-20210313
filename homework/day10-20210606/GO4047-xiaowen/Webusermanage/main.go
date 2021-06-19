package main

import (
	"net/http"
	_ "webusermanage/router"
)

func main() {

	addr := ":9999"
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/login/", http.StatusFound)
	// })
	http.ListenAndServe(addr, nil)

}
