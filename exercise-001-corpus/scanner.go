package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "sort"
)

func main() {
  wordCount := make(map[string]int)

  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanWords) 
  for scanner.Scan() {
    wordCount[scanner.Text()]++
  }
  
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  
  sort.Sort(sort.Reverse(wordCount))
  for key, value := range wordCount {
    fmt.Println(key, value)
  }
}
