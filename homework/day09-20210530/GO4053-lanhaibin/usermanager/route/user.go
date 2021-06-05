package route

import (
	"net/http"
	"usermanager/view"
)

func init() {
	http.HandleFunc("/user/create", view.HandleCreateUser)
	http.HandleFunc("/user/delete", view.HandleDeleteUser)
	http.HandleFunc("/user/query", view.HandleQueryUser)
	http.HandleFunc("/user/modify", view.HandleModifyUser)
}
