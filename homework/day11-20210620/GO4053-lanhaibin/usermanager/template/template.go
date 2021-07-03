package template

import (
	"fmt"
	htemplate "html/template"
	"io"
	"path/filepath"
)

const template_dir = "templates"

func RenderTemplate(w io.Writer, file string, data interface{}) {
	path := filepath.Join(template_dir, "**/*.html")
	// t, _ := htemplate.ParseFiles(path)
	t, err := htemplate.ParseGlob(path)
	if err != nil {
		fmt.Println(err)
	}
	// t, err = t.ParseGlob(filepath.Join(template_dir, "*.html"))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	err = t.ExecuteTemplate(w, file, data)
}
