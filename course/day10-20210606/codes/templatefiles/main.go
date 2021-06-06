package main

import (
	"os"
	"text/template"
)

func main() {
	// tpl, _ := template.New("tpl").ParseFiles("tpl/header.html", "tpl/page.html")
	// tpl, _ := template.New("tpl").ParseGlob("tpl/*.html")
	tpl, _ := template.ParseGlob("tpl/*.html")
	tpl.ExecuteTemplate(os.Stdout, "page.html", struct{ Title string }{"模板文件"})
}
