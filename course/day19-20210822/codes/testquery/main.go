package main

import (
	"log"
	"time"

	"github.com/osquery/osquery-go"
)

func main() {
	path := "/root/.osquery/shell.em"
	client, err := osquery.NewClient(path, 3*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	resp, err := client.Query("select * from system_info;")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Status.Code)
	log.Println(resp.Response)
}
