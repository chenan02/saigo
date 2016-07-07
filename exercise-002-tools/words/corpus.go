package corpus

// WordCount represents a word and its count read from the given map
type WordCount struct {
	Word  string
	Count int
}

// ByCount is used to order words by their counts
type ByCount []WordCount

// BuildWordCountSlice builds a slice for defined order of wordCounts from map
func BuildWordCountSlice(m map[string]int) []WordCount {
	var wordCounts []WordCount
	for k, v := range m {
		newWC := WordCount{k, v}
		wordCounts = append(wordCounts, newWC)
	}
	return wordCounts
}

func (bc ByCount) Len() int {
	return len(bc)
}

func (bc ByCount) Less(i, j int) bool {
	if bc[i].Count != bc[j].Count {
		return bc[i].Count > bc[j].Count
	}
	return bc[i].Word < bc[j].Word
}

func (bc ByCount) Swap(i, j int) {
	bc[i], bc[j] = bc[j], bc[i]
}
