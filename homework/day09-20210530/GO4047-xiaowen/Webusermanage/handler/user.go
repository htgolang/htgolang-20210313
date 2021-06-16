package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webusermanage/config"
	"webusermanage/user"
)
func AddHande(w http.ResponseWriter, r *http.Request) {
	u := config.UserStructure{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil{
		log.Fatalln(err)
	}
	v, err := user.Add(u)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Fprint(w, v)
}

func QueryHande(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	fmt.Println(name)
	data := user.Query(name)
	fmt.Fprint(w, data)
}

func DelHande(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	data := user.Delete(id)
	fmt.Fprint(w, data)
}

func ModifyHande(w http.ResponseWriter, r *http.Request) {
	u := config.UserStructure{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil{
		log.Fatalln(err)
	}
	r.ParseForm()
	id := r.FormValue("id")
	data := user.Modif(u, id)
	fmt.Fprint(w, data)
}

func DatilHande(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	id := r.FormValue("id")
	data := user.GetDtail(id)
	fmt.Fprint(w, data)
}