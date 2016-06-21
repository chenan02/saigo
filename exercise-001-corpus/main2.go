package main

import (
  "os"
  "fmt"
  "bufio"
  "log"
//  "sort"
)

func orderHashKeys(m map[string]int) []string {
  var keys []string
  for k := range m {
    keys = append(keys, k)
  }
  return keys
}

//func (this ByCount) Len() int {
//  return 0
//}

//func (this ByCount) Less(i, j int) bool {
//  return true
//}

//func (this ByCount) Swap(i, j int) {
//  return
//}

func main() {
  filename := os.Args[1]
  
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  wordCount := make(map[string]int)
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanWords) 
  for scanner.Scan() {
    wordCount[scanner.Text()]++
  }

  keys := orderHashKeys(wordCount)
  for _, x := range keys {
    fmt.Println(x)
  }
  
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
} 
