package main

import (
	_ "cmdb/routers"
	"net/http"
)

func main() {
	/*
		获取所有用户：
			curl --location --request GET 'http://127.0.0.1:10086/user/' --header 'User-Agent: apifox/1.0.0 (https://www.apifox.cn)'
		获取特定用户：
			curl --location --request GET 'http://127.0.0.1:10086/user/query/?id=2' --header 'User-Agent: apifox/1.0.0 (https://www.apifox.cn)'
		添加用户：
			curl --location --request POST '{{BASE_URL}} http://127.0.0.1:10086/user/add/' \
			--header 'User-Agent: apifox/1.0.0 (https://www.apifox.cn)' \
			--data-urlencode 'name=' \
			--data-urlencode 'sex=' \
			--data-urlencode 'passwd=' \
			--data-urlencode 'addr=' \
			--data-urlencode 'tel='
		删除用户：
			curl --location --request GET 'http://127.0.0.1:10086/user/delete/?id=' \
			--header 'User-Agent: apifox/1.0.0 (https://www.apifox.cn)'
		变更用户信息：
			
	*/
	addr := ":10086"
	http.ListenAndServe(addr, nil)
}
