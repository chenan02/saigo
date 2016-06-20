package main

import "fmt"

func main() {
  fmt.Println("100 fahrenheight is", fToC(100), "celsius")
  fmt.Println("100 ft is", ftToM(100), "meters")
}

func fToC(f float64) float64 {
  return (f - 32) * 5/9
}

func ftToM(ft float64) float64 {
  return ft * 0.3048
}
