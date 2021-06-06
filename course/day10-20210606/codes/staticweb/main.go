package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type StaticHandler struct {
	dir string
}

func NewStaticHandler(dir string) *StaticHandler {
	return &StaticHandler{
		dir: dir,
	}
}

func (h *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fpath := path.Join(h.dir, path.Clean(r.URL.Path))

	fmt.Println("file:", fpath)
	if fhandler, err := os.Open(fpath); err == nil {
		io.Copy(w, fhandler)
		fhandler.Close()
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func main() {
	// url => 处理器/处理器函数
	// 启动web服务
	// http.Handle("/static/", http.FileServer(http.Dir(".")))
	// /static/xxx.file => ./static/xxx.file 404
	// url规范 格式
	// protocol://username:password@addr:port/path?queryparams#fragment
	// http://addr:port/path?queryparams#fragment

	addr := ":9999"
	http.Handle("/static/", NewStaticHandler("./www/"))
	http.ListenAndServe(addr, nil)

	// socket
	/*
		Listen
		for {
			client Accept
			go handleClient(client)
		}

		handleClient(client):
			读取客户端数据 解析 => Request
			找URL和处理器的关系 => Request.URL = handler => 多路复用器 => 用于处理url和handler的关系
			调用handler.ServeHTTP函数(w, r)
			给客户端响应 => w => Client
			关闭客户端连接

		多路复用器 =>
			注册 url => handler
			程序执行过程中通过 url查找handler

		默认多路复用器:
			请求的url => 查找注册url => 找最长匹配的
			/static/ 1
			/static/www/ 2


			请求url /static/www/a.txt 匹配第二个
			请求url /static/www 匹配第一个

		go其他框架 =>
			重新定义多路复用器
			提供一些封装函数

		gin  => httprouter => 重新定义了多路复用器 => url handler查找关系 radixtree算法
			json封装

	*/
}
