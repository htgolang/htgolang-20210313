package main

import (
	"fmt"
	"net/http"
	"time"
)

type TimeHandler struct {
}

func (h *TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("time Handler")
	// r => request => 对HTTP请求封装
	fmt.Fprint(w, time.Now().Format("2006-01-02 15:04:05"))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "t2: %s", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	addr := ":9999"
	http.Handle("/t1/", &TimeHandler{})
	http.HandleFunc("/t2/", timeHandler)
	http.ListenAndServe(addr, nil)
}
