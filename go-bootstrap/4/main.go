package main

import(
    "html/template"
		"net/http"
		"log"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))

type Data struct {
    Title string // Must be exported!
    Body string  // Must be exported!
}

// Renders the templates
func renderTemplate(w http.ResponseWriter, tmpl string, page *Data) {
	err := templates.ExecuteTemplate(w, tmpl, page)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := &Data{Title:"Home page",	Body:"Welcome to our brand new home page."}
	renderTemplate(w, "index", page)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	page := &Data{Title:"About page",	Body:"This is our brand new about page."}
	renderTemplate(w, "index", page)
}


func main(){
		http.HandleFunc("/", IndexHandler)
		http.HandleFunc("/about/", AboutHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
