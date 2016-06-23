package wordcount

import (
  "testing"
  "sort"
)

type testpair struct {
  map map[string]int
  sorted []wordCount
}

standardMap := testpair{
  {"hello":3,
  "hey":3,
  "yo":1,
  "dood":2},
  []wordCount{}
}

equalMap := map[string]int{
  "yo":2,
  "ok":2,
  "hey":2
}

oneWordMap := map[string]int{
  "whatsup":5
}

emptyMap := map[string]int{
}


func TestStandard(t *testing.T) {
  wordCounts := OrderWords(standardMap)
  sort.Sort(ByCount(wordCounts))
  if wordCounts != 
}

 
