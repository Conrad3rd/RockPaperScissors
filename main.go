package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Fprintf pritn to the parameter ResponseWriter w
	html := `<h1>Hello</h1>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
