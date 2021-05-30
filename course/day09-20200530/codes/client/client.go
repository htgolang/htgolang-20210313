package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:9999/header/", nil)
	req.Header.Add("xxxxx", "yyyy")
	req.AddCookie(&http.Cookie{
		Name:  "xxxx",
		Value: "yyyyx",
	})

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Proto)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)

	io.Copy(os.Stdout, resp.Body)
}
