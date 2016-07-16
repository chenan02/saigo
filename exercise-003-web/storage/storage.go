package storage

import (
  "bufio"
  "os"
)

func WriteToFile(name string) {
  f, err := os.OpenFile("names.txt", os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  if _, err := f.WriteString(name + "\n"); err != nil {
    panic(err)
  }
}

func ReadFromFile() map[string]int {
  nc := make(map[string]int)
  f, err := os.Open("names.txt")
  if err != nil {
    panic(err)
  }
  defer f.Close()
  scanner := bufio.NewScanner(f)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
    nc[scanner.Text()]++
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }
  return nc
}

