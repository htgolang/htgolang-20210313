package utils

import (
	"html/template"
	"log"
	"net/http"
)

//解析和执行模板
func ParseTemplate(w http.ResponseWriter, r *http.Request, files []string, tplName string, data interface{}) error {
	tpl, err := template.New("tpl").ParseFiles(files...)
	if err != nil {
		log.Println("parse template file error.", err)
		return err
	}

	err = tpl.ExecuteTemplate(w, tplName, data)
	if err != nil {
		log.Println("execute template error", err)
		return err
	}
	return nil
}
