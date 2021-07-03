package route

import (
	"net/http"
	"usermanager/decorator"
	"usermanager/view"
)

func init() {
	http.HandleFunc("/department/create", decorator.LoginRequired(view.HandleCreateDepartment))
	http.HandleFunc("/department/delete", decorator.LoginRequired(view.HandleDeleteDepartment))
	http.HandleFunc("/department", decorator.LoginRequired(view.HandleQueryDepartment))
	http.HandleFunc("/department/modify", decorator.LoginRequired(view.HandleModifyDepartment))
}
