package main

import (
  "io/ioutil"
  "os"  
)

func main() {
  input := os.Args[1]
  b, err := ioutil.ReadFile(input)
  if err != nil {
    panic(err)
  }
  err = ioutil.WriteFile("output.txt", b, 0644)
  if err != nil {
    panic(err)
  } 
}
