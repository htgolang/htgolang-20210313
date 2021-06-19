package main

import (
	_ "cmdb/routers"
	"net/http"
)

func main() {
	addr := "172.24.239.212:9999"
	http.ListenAndServe(addr, nil)
}
