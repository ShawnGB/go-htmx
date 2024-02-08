//go:generate npm run build
package main

import (
	"htmx-gp/handlers"
	"log"
	"net/http"
)

func main() {

	handlers.LoadTemplates()

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", handlers.IndexHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
