package main

import (
  "html/template"
  "net/http"
  "os"
  "bufio"
)

func writeToFile(name string) {
  f, err := os.OpenFile("names.txt", os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  if _, err := f.WriteString(name + "\n"); err != nil {
    panic(err)
  }
}

func readFromFile() map[string]int {
  nc := make(map[string]int)
  f, err := os.Open("names.txt")
  if err != nil {
    panic(err)
  }
  defer f.Close()
  scanner := bufio.NewScanner(f)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
    nc[scanner.Text()]++
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }
  return nc
}

var indexT = template.Must(template.ParseFiles("index.html"))

func index(w http.ResponseWriter, r *http.Request) {
  indexT.Execute(w, nil)
}

func addName(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  name := r.Form.Get("name")
  writeToFile(name)  
  nc := readFromFile()
  indexT.Execute(w, &nc)
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/addname", addName)
  http.ListenAndServe(":8080", nil)
} 
