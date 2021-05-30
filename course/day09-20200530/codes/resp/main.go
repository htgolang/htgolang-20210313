package main

import (
	"fmt"
	"net/http"
)

func main() {
	addr := ":9999"
	http.HandleFunc("/header/", func(w http.ResponseWriter, r *http.Request) {
		// 设置状态码
		var headers http.Header = w.Header()
		headers.Add("Server", "kkserver")
		headers.Set("xxxx", "xxxxx")

		w.WriteHeader(http.StatusCreated)

		fmt.Fprintln(w, "headers")
	})

	http.HandleFunc("/redirect/", func(w http.ResponseWriter, r *http.Request) {
		// 跳转 /home/
		http.Redirect(w, r, "/home/", http.StatusFound)
		// 响应 302
		// Location: /home/
		// 重新发起请求 /home/ GET
	})

	http.HandleFunc("/home/", func(w http.ResponseWriter, r *http.Request) {
		// 跳转 /home/
		fmt.Fprintln(w, "home")
	})
	http.ListenAndServe(addr, nil)
}
