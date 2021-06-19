package handlers

import (
	"cmdb/modules"
	"cmdb/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func QueryUsersHandler(w http.ResponseWriter, r *http.Request) {
	resData := map[string]interface{}{}
	msg := ""

	//获取查询关键字
	keyword := r.FormValue("q")

	users := services.Query(keyword)

	if len(users) == 0 {
		msg = "没有匹配用户"
	}
	resData["ok"] = true
	resData["data"] = users
	resData["Msg"] = msg
	b, err := json.Marshal(resData)
	if err != nil {
		return
	}
	w.Write(b)

}

func QueryUserDetailHandler(w http.ResponseWriter, r *http.Request) {
	resData := map[string]interface{}{}
	msg := ""

	//获取用户id
	id := r.FormValue("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte("ID错误"))
		return
	}

	user := services.QueryByID(uid)

	if len(user) == 0 {
		msg = "指定ID用户不存在"
		resData["ok"] = false
		resData["Msg"] = msg
		resData["data"] = ""
	} else {
		resData["ok"] = true
		resData["data"] = user[0]
		resData["Msg"] = msg
	}

	b, err := json.Marshal(resData)
	if err != nil {
		return
	}
	w.Write(b)

}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	resData := map[string]interface{}{}
	msg := ""

	//获取提交数据
	var user modules.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		w.Write([]byte("提交数据错误，请检查"))
		return
	}

	result := services.Add(user)
	if result != "" {
		msg = result
		resData["ok"] = false
		resData["Msg"] = msg
	} else {
		resData["ok"] = true
		resData["Msg"] = msg
	}

	b, err := json.Marshal(resData)
	if err != nil {
		return
	}
	w.Write(b)
}

func DelUserHandler(w http.ResponseWriter, r *http.Request) {
	resData := map[string]interface{}{}
	msg := ""

	//获取id
	id := r.FormValue("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte("ID错误"))
		return
	}

	result := services.DelUser(uid)
	if result != "" {
		msg = result
		resData["ok"] = false
		resData["Msg"] = msg
	} else {
		resData["ok"] = true
		resData["Msg"] = msg
	}

	b, err := json.Marshal(resData)
	if err != nil {
		return
	}
	w.Write(b)
}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	resData := map[string]interface{}{}
	msg := ""

	//获取提交数据
	var user modules.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		w.Write([]byte("提交数据错误，请检查"))
		return
	}

	result := services.Edit(user)
	if result != "" {
		msg = result
		resData["ok"] = false
		resData["Msg"] = msg
	} else {
		resData["ok"] = true
		resData["Msg"] = msg
	}

	b, err := json.Marshal(resData)
	if err != nil {
		return
	}
	w.Write(b)
}
