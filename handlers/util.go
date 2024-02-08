package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var Templates *template.Template

func LoadTemplates() {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	pattern := filepath.Join(dir, "ui", "*.html")

	Templates, err = template.ParseGlob(pattern)

	if err != nil {
		panic(err)
	}
}

func RenderTemplates(w http.ResponseWriter, tmpl string, data interface{}) {

	err := Templates.ExecuteTemplate(w, tmpl+".html", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
