package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	header := `{{ define "header" }}<head><meta charset="utf-8"/><title>{{ .Title }}</title></head>{{ end }}`
	page1 := `{{ define "page1" }} <!DOCTYPE html>
	<html>
		{{ template "header" . }}
		<body>
		page1
		</body>
	</html>
	{{ end }}
	`
	page2 := `{{ define "page2" }}<!DOCTYPE html>
	<html>
		{{ template "header" . }}
		<body>
		page2
		</body>
	</html>
{{ end }}
	`

	tpl, _ := template.New("tpl").Parse(header)
	tpl, _ = tpl.Parse(page1)
	tpl, _ = tpl.Parse(page2)
	tpl.ExecuteTemplate(os.Stdout, "header", struct {
		Title string
	}{"template 定义"})
	fmt.Println()
}
