package main

import (
  "os"
  "fmt"
  "bufio"
  "log"
  "sort"
  "github.com/saigo/exercise-001-corpus/wordcount"
)

func main() {
  filename := os.Args[1]
  
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  // abstract to other libraries, raise error in main & return 
  // return err to here, raise here
  // keep as much code as possible outta here
 
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

  wordCounts := wordcount.BuildWordCountSlice(wordMap)
  sort.Sort(wordcount.ByCount(wordCounts))
  // custom wordCount printLn method in WordCount library, can make Word and Count "private"
  for _, x := range wordCounts {
    fmt.Println(x.Word, x.Count)
  }
} 
