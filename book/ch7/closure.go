package main

import "fmt"

func main() {
  total := 0
  add := func(operants ...int) {
    for _, value := range operants {
      total += value
    }
  }
  add(1, 2, 3, 4, 5)
  fmt.Println(total)
}
