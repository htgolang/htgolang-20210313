package main

import "net/http"

func main() {
	addr := ":9999"
	// 需要定义请求的URL
	// URL映射的目录
	http.Handle("/static/", http.FileServer(http.Dir("./www")))
	http.Handle("/static1/", http.FileServer(http.Dir("./www")))
	http.Handle("/static2/", http.FileServer(http.Dir("./www")))
	http.ListenAndServe(addr, nil)

	// http://addr:port/static/+path
	// ./www/+static + path

	// http://localhost:9999/static/kk.txt
	// ./www/static/kk.txt
}
