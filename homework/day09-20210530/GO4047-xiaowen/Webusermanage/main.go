package main

import (
	"fmt"
	"net/http"
	_ "webusermanage/router"
)

func main() {

	addr := ":9999"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "haha")
	})
	http.ListenAndServe(addr, nil)

}
