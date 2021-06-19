package template

import (
	htemplate "html/template"
	"io"
	"path/filepath"
)

const template_dir = "templates"

func RenderTemplate(w io.Writer, file string, data interface{}) {
	path := filepath.Join(template_dir, file)
	t, _ := htemplate.ParseFiles(path)
	t.ExecuteTemplate(w, file, data)
}
