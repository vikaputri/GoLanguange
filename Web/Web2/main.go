package main

import (
	"html/template"
	"net/http"
)

var tmplt *template.Template

func init() {
	tmplt = template.Must(template.ParseGlob("templates/*.html"))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome <h1>to first web page</h1> in Golang")
	tmplt.ExecuteTemplate(w, "index.html", nil)
}

func handlerAbout(w http.ResponseWriter, r *http.Request) {
	tmplt.ExecuteTemplate(w, "about.html", nil)
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/about", handlerAbout)
	http.ListenAndServe(":999", nil)
}
