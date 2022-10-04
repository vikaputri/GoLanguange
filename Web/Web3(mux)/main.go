package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tmplt *template.Template

func init() {
	tmplt = template.Must(template.ParseGlob("templates/*.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_ = r
	tmplt.ExecuteTemplate(w, "index.html", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmplt.ExecuteTemplate(w, "contact.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmplt.ExecuteTemplate(w, "about.html", nil)
}

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fs))
	//http.Handle("/assets/", http.StripPrefix("/assets", fs))
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/contact", contactHandler)
	r.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":800", r)

}
