package main

import (
	"fmt"
	"net/http"
)

type Mux struct {
	routers map[string]http.Handler
}

func NewMux() *Mux {
	return &Mux{
		routers: map[string]http.Handler{},
	}
}

func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	if handler, ok := mux.routers[r.URL.Path]; ok {
		handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (mux *Mux) Handle(url string, h http.Handler) {
	mux.routers[url] = h
}

func main() {

	mux := NewMux()
	// a.xxx => 调用函数
	mux.Handle("/a.xxx", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("a.xxxxx")
		fmt.Fprintln(w, "a.xxxx")
	}))

	addr := ":9999"
	http.ListenAndServe(addr, mux)
}
