package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	addr := ":9999"
	http.HandleFunc("/json/", func(w http.ResponseWriter, r *http.Request) {
		rt := map[string]string{
			"a":   "1111",
			"xxx": "xxx",
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(rt)
	})
	http.ListenAndServe(addr, nil)
}
