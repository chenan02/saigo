package main

import (
  "os"
  "fmt"
  "bufio"
  "log"
  "sort"
)

type WordCount struct {
  Word string
  Count int
}

type ByCount []WordCount

func orderWordCounts(m map[string]int) []WordCount {
  var wordCountSlice []WordCount 
  for k, v := range m {
    newWC := WordCount{k, v}
    wordCountSlice = append(wordCountSlice, newWC)
  }
  return wordCountSlice
}

func (this ByCount) Len() int {
  return len(this)
}

func (this ByCount) Less(i, j int) bool {
  return this[i].Count > this[j].Count
}

func (this ByCount) Swap(i, j int) {
  this[i], this[j] = this[j], this[i]
}

func main() {
  filename := os.Args[1]
  
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  wordMap := make(map[string]int)
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanWords) 
  for scanner.Scan() {
    wordMap[scanner.Text()]++
  }

  wordCountSlice := orderWordCounts(wordMap)
  sort.Sort(ByCount(wordCountSlice))
  for _, x := range wordCountSlice {
    fmt.Println(x.Word, x.Count)
  }
  
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
} 
