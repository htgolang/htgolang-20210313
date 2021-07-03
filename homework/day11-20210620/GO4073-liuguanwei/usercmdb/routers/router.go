package routers

import (
	"net/http"
	"usercmdb/controllers"
)

func init() {
	http.HandleFunc("/", controllers.GetUsers)
	http.HandleFunc("/create", controllers.AddUser)
	http.HandleFunc("/delete", controllers.DeleteUser)
	http.HandleFunc("/modify", controllers.ModifyUser)
}
