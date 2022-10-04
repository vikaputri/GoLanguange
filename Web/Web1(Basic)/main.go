package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to first web page in Golang %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":999", nil)
}
