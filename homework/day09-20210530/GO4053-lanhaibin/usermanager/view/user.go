package view

import (
	"encoding/json"
	"net/http"
	"strconv"
	"usermanager/service"
)

type response map[string]interface{}
type request map[string]string

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	resp := make(response)
	w.Header().Set("Content-Type", "Application/json")

	var req request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	encoder := json.NewEncoder(w)
	if err != nil {
		resp["ok"] = false
		encoder.Encode(resp)
		return
	}

	name, ok := req["name"]
	if !ok {
		resp["ok"] = false
	}

	addr, ok := req["addr"]
	if !ok {
		resp["ok"] = false
	}

	tel, ok := req["tel"]
	if !ok {
		resp["ok"] = false
	}

	password, ok := req["password"]
	if !ok {
		resp["ok"] = false
	}

	if _, ok := resp["ok"]; ok {
		encoder.Encode(resp)
		return
	}

	err = service.Us.Create(name, addr, tel, password)
	if err != nil {
		resp["ok"] = false
	} else {
		resp["ok"] = true
	}
	encoder.Encode(resp)
}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	resp := make(response)
	encoder := json.NewEncoder(w)

	r.ParseForm()
	ids := r.FormValue("id")
	id, _ := strconv.Atoi(ids)
	err := service.Us.Delete(id)
	if err != nil {
		resp["ok"] = false
	} else {
		resp["ok"] = true
	}
	encoder.Encode(resp)
}

func HandleQueryUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	resp := make(response)
	encoder := json.NewEncoder(w)

	r.ParseForm()
	q := r.FormValue("q")
	qs := service.Us.Query(q)
	resp["ok"] = true
	resp["data"] = qs
	encoder.Encode(resp)
}

func HandleModifyUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	resp := make(response)
	var req request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {

	}
	encoder := json.NewEncoder(w)
	ids := req["id"]
	id, _ := strconv.Atoi(ids)
	name := req["name"]
	addr := req["addr"]
	tel := req["tel"]

	err = service.Us.Modify(id, name, addr, tel)
	if err != nil {
		resp["ok"] = false
	} else {
		resp["ok"] = true
	}
	encoder.Encode(w)

}
