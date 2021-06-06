package api

import (
	"encoding/json"
	"fmt"
	"mgr/model/user"
	"net/http"
)
var Users []*user.User
type ResponseJson struct {
	Ok bool `json:"ok"`
	Data []*user.User `json:"data,omitempty"`
	Msg string `json:"msg,omitempty"`
}

func MyMarshalError(w *http.ResponseWriter,msg string)  {
	var uRes ResponseJson
	uRes.Ok = false
	uRes.Msg = msg
	bytes, _ := json.Marshal(uRes)
	fmt.Fprintln(*w, string(bytes))
}

func ApiEntrance() {
	addr := ":9999"
	var responseJson = new(ResponseJson)
	//查询用户接口
	QueryApi(responseJson)
	/*
	测试:
	(1)查询所有用户的数据
	127.0.0.1:9999/query?all=""
	(2)通过id查询用户
	127.0.0.1:9999/query?id=aba  错误
	127.0.0.1:9999/query?id=1  存在
	127.0.0.1:9999/query?id=2  不存在
	*/
	//添加用户接口
	AddUserApi(responseJson)
	http.ListenAndServe(addr, nil)
}
