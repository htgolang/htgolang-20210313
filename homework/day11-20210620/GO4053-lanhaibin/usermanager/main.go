package main

import (
	"net/http"
	_ "usermanager/route"
)

func main() {
	// path := "db/user.json"
	// p := persistent.JsonPersistent{DbPath: path}
	// p.Decode(&model.Users)
	// go func() {
	// 	for {
	// 		time.Sleep(3 * time.Second)
	// 		p.Encode(&model.Users)
	// 	}
	// }()

	addr := ":8080"
	http.ListenAndServe(addr, nil)
}
