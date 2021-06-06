package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	url := "http://localhost:9999/form-data/"

	buffer := bytes.NewBufferString("")

	writer := multipart.NewWriter(buffer)
	file, _ := writer.CreateFormField("xxx")
	ctx, _ := ioutil.ReadFile("a.xxxx")
	file.Write(ctx)

	writer.Close()
	// fmt.Println(writer.FormDataContentType())
	// fmt.Println(buffer.WriteTo(os.Stdout))

	resp, err := http.Post(url, writer.FormDataContentType(), buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Proto)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)

	io.Copy(os.Stdout, resp.Body)
}
