package main

import "fmt"

func sum(numbers []int) int {
  total := 0
  for _, value := range numbers {
    total += value
  }
  return total
}

func half(number int) (int, bool) {
  return number/2, number%2 == 0
}

func greatest(args ...int) int {
  biggest := args[0]
  for _, value := range args {
    if biggest < value {
      biggest = value
    }
  }
  return biggest
}

func makeOddGenerator() func() uint {
  i := uint(1)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

func fib(n int) int {
  if n == 0 {
    return 0
  }
  if n == 1 {
    return 1
  }
  return fib(n-1) + fib(n-2)
}

func deferring() {
  str := recover()
  fmt.Println(str)
 }

func main() {
  numbers := []int{1, 2, 3, 4, 5}
  total := sum(numbers)
  fmt.Println("sum:", total)

  halved, isEven := half(2)
  fmt.Println("half:", halved, isEven)

  biggest := greatest(2,4,1,6,3)
  fmt.Println("greatest:", biggest)

  nextOdd := makeOddGenerator()
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())

  fib5 := fib(5)
  fmt.Println("Fib:", fib5)

  defer deferring()
  panic("YO IM PANICKING!!")
}
