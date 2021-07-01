package utils

import (
	"net/http"
	"text/template"
)

func ParseTemplate(w http.ResponseWriter, r *http.Request,
	files []string, tplName string,
	data interface{}) {
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(w, tplName, data)
	if err != nil {
		panic(err)
	}
}
