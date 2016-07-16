package main

import (
  "html/template"
  "net/http"
)

var indexT = template.Must(template.ParseFiles("index.html"))
var nc = make(map[string]int)

func index(w http.ResponseWriter, r *http.Request) {
  indexT.Execute(w, nil)
}

func addName(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  name := r.Form.Get("name")
  nc[name]++  
  indexT.Execute(w, &nc)
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/addname", addName)
  http.ListenAndServe(":8080", nil)
} 
