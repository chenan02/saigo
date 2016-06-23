package wordcount

type WordCount struct {
  Word string
  Count int
}

type ByCount []WordCount

// Generate slice for defined order of wordCounts from map
func BuildWordCountSlice(m map[string]int) []WordCount {
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
