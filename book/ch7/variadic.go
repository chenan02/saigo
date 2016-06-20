package main

import "fmt"

func main() {
  total := add(1, 2, 3, 4, 5)
  fmt.Println(total)
}

func add(args ...int) int {
  total := 0
  for _, value := range args {
    total += value
  }
  return total
}
