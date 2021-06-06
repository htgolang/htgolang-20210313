package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	params := url.Values{}
	params.Add("a", "1")
	params.Add("a", "2")
	params.Set("b", "3")
	params.Set("b", "4")
	url := "http://localhost:9999/form/"
	fmt.Println(url)
	resp, err := http.PostForm(url, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Proto)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)

	io.Copy(os.Stdout, resp.Body)

}
