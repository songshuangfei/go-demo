package handler

import (
	"html/template"
	"net/http"
	"os"
)

const (
	uploadDir = "./uploads/"
	viewDir   = "./view/"
)

func isExists(path string) bool { //判断问价你是否存在
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func renderHTML(w http.ResponseWriter, viewname string, renderData map[string]interface{}) error {
	t, err := template.ParseFiles(viewDir + viewname + ".html")
	if err != nil {
		return nil
	}
	err = t.Execute(w, renderData)
	return err
}
