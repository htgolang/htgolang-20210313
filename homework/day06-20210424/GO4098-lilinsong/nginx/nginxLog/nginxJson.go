package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"time"
)

type NginxJsonLog struct {
	Timestamp     *time.Time `json:"@timestamp"`
	Host          string     `json:"host"`
	ClientIP      string     `json:"clientip"`
	Size          int        `json:"size"`
	ResponseTime  float32    `json:"responsetime"`
	UpstreamTime  string     `json:"upstreamtime"`
	UpstreamHost  string     `json:"upstreamhost"`
	HttpHost      string     `json:"http_host"`
	Url           string     `json:"url"`
	Domain        string     `json:"domain"`
	Xff           string     `json:"xff,omitempty"`
	Referer       string     `json:"referer,omitempty"`
	TcpXff        string     `json:"tcp_xff,omitempty"`
	HttpUserAgent string     `json:"http_user_agent"`
	Status        string     `json:"status"`
}

func nginxLog(file string) error {
	var nginx = new(NginxJsonLog)
	var request = make(map[string]map[string]int)
	fileObj, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)

	for {
		lien, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if err := json.Unmarshal(lien, &nginx); err != nil {
			fmt.Println(err)
		}

		if _, ok := request[nginx.ClientIP]; !ok {
			request[nginx.ClientIP] = make(map[string]int)
		}
		if _, ok := request[nginx.ClientIP][nginx.Status]; !ok {
			request[nginx.ClientIP][nginx.ClientIP]++
			request[nginx.ClientIP][nginx.Status]++
		} else {
			request[nginx.ClientIP][nginx.ClientIP]++
			request[nginx.ClientIP][nginx.Status]++
		}
	}

	var tmp = make([][]string, len(request))
	index := 0
	for key, value := range request {
		tmp[index] = append(tmp[index], key, strconv.Itoa(value[key]), strconv.Itoa(value["200"]), strconv.Itoa(value["400"]))
		index++
	}
	sort.Slice(tmp, func(i, j int) bool {
		a, _ := strconv.Atoi(tmp[i][1])
		b, _ := strconv.Atoi(tmp[j][1])
		return a > b
	})
	for i := 0; i < 10; i++ {
		fmt.Printf("IP:%s, 访问次数:[%s], 状态码200:[%s], 状态码404:[%s]\n", tmp[i][0], tmp[i][1], tmp[i][2], tmp[i][3])
	}
	return nil
}

func main() {
	nginxLog("./test.log")
}
