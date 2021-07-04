package handlers

import (
	"cmdb/web"
	"fmt"
	"net/http"
)

type WebHandler func(int64, http.ResponseWriter, *http.Request)

func AuthWrapper(handler WebHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("auth wrapper before")
		// 验证用户是否登陆
		token := web.LoadToken(w, r) // 加载token
		uid, ok := token.Get("uid").(int64)

		if !ok {
			// 已登录, 跳转到登陆页面
			http.Redirect(w, r, "/login/", http.StatusFound)
			return
		}
		fmt.Println("auth wrapper after: %d", uid)
		// 验证成功(已登陆)
		handler(uid, w, r)
	}
}
