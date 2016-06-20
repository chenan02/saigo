package main

import "fmt"

func main() {
  fmt.Printf("Please enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)
  
  output := input * 2

  fmt.Println(output)
}

// x := 5, var x int = 5
// 6
// a var lives within the nearest curly braces and whats nested inside
// cant change const dawggg  
