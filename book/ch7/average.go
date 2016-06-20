package main

import "fmt"

func main() {
  x := []float64{1, 2, 3, 4, 5}
  average := average(x)
  fmt.Println(average)  
}

func average(x []float64) float64 {
  total := 0.0
  for _, value := range x {
    total += value
  }
  return total / float64(len(x))
}
