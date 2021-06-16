package router

import (
	"net/http"
	"webusermanage/handler"
)


func init() {
	http.HandleFunc("/user/add/", handler.AddHande)
	http.HandleFunc("/user/query/", handler.QueryHande)
	http.HandleFunc("/user/delete/", handler.DelHande)
	http.HandleFunc("/user/modify/", handler.ModifyHande)
	http.HandleFunc("/user/datil/",handler.DatilHande)
}