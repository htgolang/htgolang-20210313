package handlers

import (
	"encoding/base64"
	"net/http"
	"strings"

	"exporter/services"
)

func AuthHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检查前
		authorization := r.Header.Get("Authorization")
		prefix := "Basic "
		if strings.HasPrefix(authorization, prefix) {
			authorization = authorization[len(prefix):]
		}
		txt, err := base64.StdEncoding.DecodeString(authorization)
		if err != nil || !services.ValidateUser(string(txt)) {
			// 认证失败
			w.Header().Add("www-authenticate", "Basic")
			w.WriteHeader(401)
			return
		}

		handler.ServeHTTP(w, r)
		// 检查后
	})
}
