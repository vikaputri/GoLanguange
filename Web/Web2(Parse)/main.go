package main

import (
	"html/template"
	"net/http"
)

var tmplt *template.Template

func init() {
	tmplt = template.Must(template.ParseGlob("templates/*.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmplt.ExecuteTemplate(w, "index.html", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmplt.ExecuteTemplate(w, "contact.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmplt.ExecuteTemplate(w, "about.html", nil)
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":999", nil)
}
