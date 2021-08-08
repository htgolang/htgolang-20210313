package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/imroc/req"
)

func main() {
	req.Debug = true
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	request := req.New()
	// 心跳
	// {uuid: "xx"}
	response, err := request.Post("https://localhost:8888/v1/agent/heartbeat", req.Param{
		"uuid": "xxxx",
	}, req.Header{"x-token": "820923a1a2dad74e8d279c48b8a1160c"}, client)

	if err == nil {
		rt := make(map[string]interface{})
		response.ToJSON(&rt)
		fmt.Println(rt)
	}
}
