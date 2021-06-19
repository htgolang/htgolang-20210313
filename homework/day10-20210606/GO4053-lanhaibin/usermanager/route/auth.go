package route

import (
	"net/http"
	"usermanager/decorator"
	"usermanager/view"
)

func init() {
	http.HandleFunc("/login", view.HandleLogin)
	http.HandleFunc("/logout", decorator.LoginRequired(view.HandleLogout))
}
