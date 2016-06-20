package main

import "fmt"

func main() {
  array := [5]int{1, 2, 3, 4, 5}
//  slice := []int{1, 2, 3, 4, 5}
//  hash := map[string]int {"one":1, "two":2, "three":3}
  
  fourth := array[3]
  fmt.Println("fourth:", fourth)
 
  //2nd param = size slice, 3rd = size underlying array 
  slice2 := make([]int, 3, 9) 
  fmt.Println("slice2:", slice2)
  fmt.Println("slice2 size:", len(slice2))
  
  x := [6]string{"a","b","c","d","e","f"}
  twotofive := x[2:5]
  fmt.Println("twotofive:", twotofive)
}
