package handlers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func AuthHandler(handler http.Handler, user, password string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检查前
		authorization := r.Header.Get("Authorization")
		prefix := "Basic "
		if strings.HasPrefix(authorization, prefix) {
			authorization = authorization[len(prefix):]
		}
		txt, err := base64.StdEncoding.DecodeString(authorization)
		if err != nil || !validateUser(string(txt), user, password) {
			// 认证失败
			w.Header().Add("www-authenticate", "Basic")
			w.WriteHeader(401)
			return
		}

		handler.ServeHTTP(w, r)
		// 检查后
	})
}

func validateUser(authorization string, user, password string) bool {
	if user == "" && password == "" {
		// 不需要输入用户认证
		return true
	}

	pos := strings.Index(authorization, ":")
	if pos < 0 {
		return false
	}

	inputUser, inputPassword := authorization[:pos], authorization[pos+1:]

	if inputUser == user {
		return password == fmt.Sprintf("%X", md5.Sum([]byte(inputPassword)))
	}
	return false
}
