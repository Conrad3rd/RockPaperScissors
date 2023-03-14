package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}

}

func main() {
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
