package main

import "fmt"

func main() {
  elements := map[string]map[string]string {
    "H": map[string]string {
      "name":"Hydrogen",
      "state":"gas",
    },
    "He": map[string]string {
      "name":"Helium",
      "state":"gas",
    },
  }

  if element, ok := elements["H"]; ok {
    fmt.Println(element["name"], element["state"])
  }
}
