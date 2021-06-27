package router

import (
	"net/http"
	"webusermanage/handler"
)


func init() {
	http.HandleFunc("/login/", handler.LoginHande)
	http.HandleFunc("/logout/", handler.LogoutHande)
	http.HandleFunc("/user/", handler.UserHande)
	http.HandleFunc("/user/add/", handler.AddHande)
	http.HandleFunc("/user/role/", handler.RoleHande)
	http.HandleFunc("/user/delete/", handler.DelHande)
	http.HandleFunc("/user/modify/", handler.ModifyHande)
	http.HandleFunc("/user/datil/",handler.DatilHande)
}