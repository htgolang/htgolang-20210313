package api

import (
	"encoding/json"
	"fmt"
	ctlUser "mgr/controller/user"
	"net/http"
	"strconv"
	"strings"
)

func FindUserById(res *ResponseJson,id int,w http.ResponseWriter)  {
	var uRes ResponseJson
	flag:=false
	res.Data = ctlUser.Create("json").Read()
	for i := 0; i < len(res.Data); i++ {
		if id==res.Data[i].Id{
			flag = true
			uRes.Ok = true
			uRes.Data = append(uRes.Data,res.Data[i])
			encoder := json.NewEncoder(w)
			encoder.Encode(uRes)
			break
		}
	}
	if !flag{
		MyMarshalError(&w, "您搜索的用户不存在!")
	}
}

func QueryApi(res *ResponseJson) {
	http.HandleFunc("/query/", func(w http.ResponseWriter, r *http.Request) {
		var rMap  = make(map[string][]string)
		r.ParseForm() // 解析提交数据
		rMap = r.Form
		for k,_ := range rMap{
			fmt.Println("all" == k)
			if "all" == k &&""==strings.Trim(r.Form.Get("all"), "\""){
				fmt.Println("sdfsfsd")
				res.Data = ctlUser.Create("json").Read()
				if len(res.Data) >0{
					encoder := json.NewEncoder(w)
					res.Ok = true
					encoder.Encode(res)
				}else {
					MyMarshalError(&w, "查询错误,不存在用户!")
				}
			}else if "id" ==k{
				id,err := strconv.ParseInt(r.Form.Get("id"),10,32)
				if err != nil{
					MyMarshalError(&w, "请检查传入id,id必须为整数!")
					return
				}
				FindUserById(res,int(id),w)
			}else {
				MyMarshalError(&w, "请检查传入接口,请求非法!")
				return
			}
		}
	})
}