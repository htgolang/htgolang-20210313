package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "http://localhost:9999/body/"
	contentType := "xxxxx"
	buffer := bytes.NewBufferString("")

	params := make(map[string]string)
	params["name"] = "kk"
	params["xxx"] = "zz"
	params["yyy"] = "1111"

	encoder := json.NewEncoder(buffer)
	encoder.Encode(params)

	resp, err := http.Post(url, contentType, buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Proto)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)

	io.Copy(os.Stdout, resp.Body)
}
