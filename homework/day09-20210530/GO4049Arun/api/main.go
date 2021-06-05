package api

import (
	"encoding/json"
	"mgr/model/user"
	"net/http"
)
var Users []*user.User
type ResponseJson struct {
	Ok bool `json:"ok"`
	Data []*user.User `json:"data"`
	Msg string `json:"msg"`
}

func MyMarshal(uRes *ResponseJson) string {
	bytes, _ := json.Marshal(uRes)
	return string(bytes)
}

func ApiEntrance() {
	addr := ":9999"
	var responseJson = new(ResponseJson)
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

	http.ListenAndServe(addr, nil)
}
