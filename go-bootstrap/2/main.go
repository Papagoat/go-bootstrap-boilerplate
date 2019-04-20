package main

import (
	"html/template"
	"log"
	"net/http"
)

type Data struct {
    Title string // Must be exported!
		Body string // Must be exported!
}

// Handler for home page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
  page := &Data{Title:"Home page", Body:"This is the home page."}
	t, _ := template.ParseFiles("templates/index.html")
  t.Execute(w, page)
}

// Handler for about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
  page := &Data{Title:"About page", Body:"This is the about page."}
	t, _ := template.ParseFiles("templates/index.html")
  t.Execute(w, page)
}

// Main function
func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about/", AboutHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
