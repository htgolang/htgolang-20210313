package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// session cookie
	// 会话跟踪

	// http 无状态

	// 用户第一次方式生成唯一标识 (tid(session id => sid) => 服务端对应存储 => session)
	// 告知浏览器你的标识 tid Set-Cookie
	// 以后请求时携带tid Cookie tid => 找存储 => 增、删、改、差

	// cookie
	// 浏览器端存储
	// 服务器如何告知浏览器存储某些数据
	// 响应头 Set-Cookie: k=v attr (域名 path 过期时间 httponly isSecure ....)

	// 浏览器再请求中会将这些数据携带
	// Cookie: k=v
	addr := ":9999"
	http.HandleFunc("/set/", func(w http.ResponseWriter, r *http.Request) {
		cookie := http.Cookie{
			Name:  "unixtime",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
			Path:  "/",
		}
		http.SetCookie(w, &cookie)
	})

	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("----------------")
		cookies := r.Cookies()
		fmt.Println(cookies)
		unixtime, err := r.Cookie("unixtime")
		fmt.Println(unixtime, err)
		tid, err := r.Cookie("tid")
		fmt.Println(tid, err)

	})
	http.ListenAndServe(addr, nil)
}
