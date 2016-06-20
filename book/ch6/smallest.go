package main

import "fmt"

func main() {
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }
  lowest, err := findLowest(x)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println(lowest)
}

func findLowest(x []int) (int, error) {
  if len(x) == 0 {
    return nil, fmt.Errorf("bad")
  }
  lowest := x[0]
  for _, value := range x {
    if lowest > value {
      lowest = value
    }
  }
  return lowest, nil
}
  
