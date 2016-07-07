package scanner

import (
  "bufio"
  "os"
)

func ReadMap(filename string) (map[string]int, error) {
  file, err := os.Open(filename)
  if err != nil {
    return nil, err
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
    return nil, err
  }

  return wordMap, nil
}
