package main

import "fmt"

var i string = "hey"

func main() {
  print()
  i = "hello"
  fmt.Println(i)
  print()
}

func print() {
  fmt.Println(i)
}
