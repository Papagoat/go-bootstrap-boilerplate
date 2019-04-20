package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler for home page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
  title, body := "Home", "This is the home page."
	fmt.Fprintf(w, "<a href='/'>Home</a> <a href='/about/'>About</a><br><h1>%s</h1><div>%s</div>", title, body)
}

// Handler for about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
  title, body := "About", "This is the about page."
	fmt.Fprintf(w, "<a href='/'>Home</a> <a href='/about/'>About</a><br><h1>%s</h1><div>%s</div>", title, body)
}

// Main function
func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about/", AboutHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
