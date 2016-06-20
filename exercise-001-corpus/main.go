package main

import (
  "fmt"
  "io"
//  "sort"
//  "strings"
  "os"
)

func main() {
  input := os.Args[1]
  fi, err := os.Open(input)
  
  // check input ok
  if err != nil {
    panic(err)
  }

  // callback to close file/panic if file close went bad
  defer func() {
    if err := fi.Close(); err != nil {
      panic(err)
    }
  }()

  // buffer to read 1024 bytes at a time
  buf := make([]byte, 1024)
  for {
    n, err := fi.Read(buf)
  
    // panic if read goes bad
    if err != nil && err != io.EOF {  
      panic(err)
    }

    if n == 0 {
      break
    }

    fmt.Println(buf[:n])
  }
}
