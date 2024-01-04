//go:generate npm run build
package main

import (
	"htmx-gp/handlers"
	"log"
	"net/http"
)

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.Base)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
