package main

import "fmt"

func main() {
  x := 5
  xPtr := &x
  fmt.Println(xPtr)
  newPtr := new(int)
  fmt.Println(newPtr)
  
  y := 1.5
  square(&y)
  fmt.Println(y)

  a, b := 1, 2
  swap(&a, &b)
  fmt.Println(a, b)
}

func square(x *float64) {
  *x = *x * *x
}

func swap(x *int, y *int) {
  temp := *x
  *x = *y
  *y = temp
}  
