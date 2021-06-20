package main

import (
	"embed"
	"fmt"
	"html/template"
	"os"
)

//go:embed tpl
var fs embed.FS

func main() {
	login, err := fs.ReadFile("tpl/login.html")
	fmt.Println(string(login), err)
	tpl, err := template.New("").Parse(string(login))
	tpl.Execute(os.Stdout, "kk")
	fmt.Println()

	// ParseFiles ParseGlob
	templ, err := template.ParseFS(fs, "tpl/*.html")
	fmt.Println(err, templ)
	templ.ExecuteTemplate(os.Stdout, "login.html", "ada")
	fmt.Println()
	templ, err = template.ParseFS(fs, "tpl/users/*.html")
	templ.ExecuteTemplate(os.Stdout, "index.html", nil)
	fmt.Println()

	templ.ExecuteTemplate(os.Stdout, "add.html", nil)
	fmt.Println()

	templ, err = template.ParseFS(fs, "tpl/hosts/*.html")
	templ.ExecuteTemplate(os.Stdout, "index.html", nil)
	fmt.Println()
}
