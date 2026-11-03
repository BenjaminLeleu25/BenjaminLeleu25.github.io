package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	// Load our HTML template.
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Handle requests to the homepage.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Only serve the homepage at "/".
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Unable to load page", http.StatusInternalServerError)
			log.Println("Template error:", err)
		}
	})

	// Serve CSS, JavaScript and images.
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// Start the server.
	log.Println("Portfolio running at http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
