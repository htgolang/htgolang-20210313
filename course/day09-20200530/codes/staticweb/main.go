package main

import "net/http"

func main() {
	addr := ":9999"
	// 需要定义请求的URL
	// URL映射的目录
	http.Handle("/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("./www")),
		),
	)

	http.ListenAndServe(addr, nil)

	// http://addr:port/static/+path
	// ./www/ + path

	// http://localhost:9999/static/kk.txt
	// ./www/kk.txt
}
