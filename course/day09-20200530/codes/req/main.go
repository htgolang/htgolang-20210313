package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	addr := ":9999"
	http.HandleFunc("/header/", func(w http.ResponseWriter, r *http.Request) {
		// http request 第一行
		// Method URL Protocol
		fmt.Println(strings.Repeat("-", 10))
		fmt.Println("method:", r.Method)
		fmt.Println("url: ", r.URL)
		fmt.Println("proto: ", r.Proto)
		fmt.Println("host: ", r.Host)
		// 请求头 key: value
		for k, v := range r.Header {
			fmt.Println(k, v)
		}
	})

	http.HandleFunc("/queryparams/", func(w http.ResponseWriter, r *http.Request) {
		// queryparams -> r.Form
		r.ParseForm() // 解析提交数据
		fmt.Println(r.Form)
		fmt.Println(r.Form.Get("a"))
		fmt.Println(r.Form.Get("b"))
		fmt.Println(r.Form.Get("c"))
		fmt.Println(r.Form.Get("d"))
		fmt.Println(r.Form["a"])
		fmt.Println(r.Form["d"])

		fmt.Fprintln(w, time.Now())
	})

	http.HandleFunc("/queryparams2/", func(w http.ResponseWriter, r *http.Request) {
		// queryparams -> r.Form
		fmt.Println(r.FormValue("a"))
		fmt.Println(r.FormValue("b"))
		fmt.Println(r.FormValue("c"))
	})

	http.HandleFunc("/form/", func(w http.ResponseWriter, r *http.Request) {
		// queryparams -> r.Form
		r.ParseForm()
		fmt.Println(r.PostForm)
		fmt.Println(r.Form)
	})

	http.HandleFunc("/form2/", func(w http.ResponseWriter, r *http.Request) {
		// queryparams -> r.Form
		fmt.Println(r.PostFormValue("a"))
		fmt.Println(r.PostFormValue("b"))
		fmt.Println(r.PostFormValue("c"))
		fmt.Println(r.FormValue("a"))
		fmt.Println(r.FormValue("b"))
		fmt.Println(r.FormValue("x"))
		fmt.Println(r.FormValue("y"))
	})

	http.HandleFunc("/form-data/", func(w http.ResponseWriter, r *http.Request) {
		// 上传文件
		fmt.Println(r.Method)
		fmt.Println(r.Header)
		r.ParseMultipartForm(10 * 1024 * 1024)
		fmt.Println(r.MultipartForm.File)
		fmt.Println(r.MultipartForm.Value)
		fmt.Println(r.Form)
		fmt.Println(r.PostForm)
		for k, fileHeaders := range r.MultipartForm.File {
			fmt.Println("file ", k)
			for idx, fileHeader := range fileHeaders {
				fmt.Println(idx, fileHeader.Filename, fileHeader.Header, fileHeader.Size)
				if file, err := fileHeader.Open(); err == nil {
					defer file.Close()
					io.Copy(os.Stdout, file)
					fmt.Println()
				}
			}
		}
	})

	http.HandleFunc("/form-data2/", func(w http.ResponseWriter, r *http.Request) {
		// 上传文件
		f, fh, _ := r.FormFile("a")
		fmt.Println(fh.Filename, fh.Header, fh.Size)
		io.Copy(os.Stdout, f)
		fmt.Println()
		f, fh, _ = r.FormFile("b")
		fmt.Println(fh.Filename, fh.Header, fh.Size)
		io.Copy(os.Stdout, f)
		fmt.Println()
	})

	http.HandleFunc("/body/", func(w http.ResponseWriter, r *http.Request) {
		// 上传文件
		fmt.Println(strings.Repeat("-", 10))
		fmt.Println(r.Header.Get("Content-Type"))
		//io.Copy(os.Stdout, r.Body)
		req := map[string]string{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		fmt.Println(err, req)
		fmt.Println()
	})

	http.ListenAndServe(addr, nil)
}
