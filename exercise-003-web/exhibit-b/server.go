package main

import (
	"html/template"
	"net/http"
)
// template = sanitize html
var homeT = template.Must(template.ParseFiles("exhibit-b/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func main() {
  // calls home to apply the homeT template(html) to output at home
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
