package main

import (
	_ "cmdb/routers"
	"net/http"
)

func main() {
	http.ListenAndServe("172.24.239.212:9999", nil)
}
