package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("templates/base.gohtml"))

		templ.Execute(w, nil)
	}

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
