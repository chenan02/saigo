package main

import (
  "html/template"
  "net/http"
  "github.com/saigo/exercise-003-web/storage"
)

var indexT = template.Must(template.ParseFiles("index.html"))

func index(w http.ResponseWriter, r *http.Request) {
  indexT.Execute(w, nil)
}

func addName(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  name := r.Form.Get("name")
  storage.WriteToFile(name)  
  nc := storage.ReadFromFile()
  indexT.Execute(w, &nc)
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/addname", addName)
  http.ListenAndServe(":8080", nil)
} 
