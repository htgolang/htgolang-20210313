package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	ID   string
	Name string
}

func main() {
	// 重写模板中某块的内容 为内容定义一个逻辑块 {{block "name" . }} xxx {{ end }}

	// `我是{{ . }}`
	txt := `我是{{ block "desc" . }} {{ . }} {{ end }}`
	data := User{"xxx", "yyyy"}

	tpl, err := template.New("overwrite").Parse(txt)
	if err != nil {
		log.Fatal(err)
	}

	// tpl.Execute(os.Stdout, data)
	// fmt.Println()

	// 我是{{ . }},欢迎您
	overwirteTxt := `{{ define "desc" }} {{ . }},欢迎您 {{ end }}`
	tpl, err = tpl.Parse(overwirteTxt)
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(os.Stdout, data)
}
