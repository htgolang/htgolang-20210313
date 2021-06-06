package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	header := `{{ define "header" }}<head><meta charset="utf-8"/><title>{{ .Title }}</title></head>{{ end }}`
	page1 := `<!DOCTYPE html>
	<html>
		{{ template "header" . }}
		<body>
		page1
		</body>
	</html>
	`
	page2 := `<!DOCTYPE html>
	<html>
		{{ template "header" . }}
		<body>
		page2
		</body>
	</html>
	`

	tpl, _ := template.New("tpl").Parse(header)
	tpl, _ = tpl.Parse(page1)
	tpl.Execute(os.Stdout, struct {
		Title string
	}{"page 1template 定义"})
	fmt.Println()

	tpl2, _ := template.New("tpl").Parse(header)
	tpl2, _ = tpl2.Parse(page2)
	tpl2.Execute(os.Stdout, struct {
		Title string
	}{"page 2 template 定义"})

}
