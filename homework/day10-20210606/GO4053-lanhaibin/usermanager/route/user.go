package route

import (
	"net/http"
	"usermanager/decorator"
	"usermanager/view"
)

func init() {
	http.HandleFunc("/user/create", decorator.LoginRequired(view.HandleCreateUser))
	http.HandleFunc("/user/delete", decorator.LoginRequired(view.HandleDeleteUser))
	http.HandleFunc("/user", decorator.LoginRequired(view.HandleQueryUser))
	http.HandleFunc("/user/modify", decorator.LoginRequired(view.HandleModifyUser))
	// http.HandleFunc("/user", view.HandleUserList)
}
