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

// Generate slice for defined order of wordCounts from map
func orderWords(m map[string]int) []WordCount {
  var wordCounts []WordCount
  for k, v := range m {
    newWC := WordCount{k, v}
    wordCounts = append(wordCounts, newWC)
  }
  return wordCounts
}

func (this ByCount) Len() int {
  return len(this)
}

func (this ByCount) Less(i, j int) bool {
  if this[i].Count != this[j].Count {
    return this[i].Count > this[j].Count
  }
  return this[i].Word < this[j].Word
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
  
  // Read counts of words into map
  wordMap := make(map[string]int)
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanWords) 
  for scanner.Scan() {
    wordMap[scanner.Text()]++
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  wordCounts := orderWords(wordMap)
  sort.Sort(ByCount(wordCounts))
  for _, x := range wordCounts {
    fmt.Println(x.Word, x.Count)
  }
} 
