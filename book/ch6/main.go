package main

import "fmt"

func main() {
  var y [5]int
  y[4] = 100
  fmt.Println(y)
  x := [5]float64{ 98, 93, 77, 82, 83 }
  var total float64 = 0
  for _, value := range x {
    total += value
  }
  fmt.Println(total/float64(len(x)))
  
  slice := x[:]
  slice = append(slice, 6, 6, 6)
  fmt.Println(slice)
  short := make([]float64, 2)
  copy(short, slice)
  fmt.Println(short)
  
  z := make(map[string]int)
  z["key"] = 10
  if name, ok := z["key"]; ok {
    fmt.Println(name)
  }
  if name, ok := z["yo"]; ok {
    fmt.Println(name)
  }
}

