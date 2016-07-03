package wordcount

import (
  "testing"
  "sort"
  "github.com/stretchr/testify/assert"
)

func TestStandard(t *testing.T) {
  expectedSort := []WordCount{{"hello",3}, {"hey",3}, {"dood",2}, {"yo",1}}
  standardMap := map[string]int{"hey":3, "yo":1, "dood":2, "hello":3}
  wordCounts := BuildWordCountSlice(standardMap)
  if assert.NotNil(t, wordCounts) {
    assert.Equal(t, len(wordCounts), 4, "The WordCount slice should have correct size")
    sort.Sort(ByCount(wordCounts))
    assert.Equal(t, wordCounts, expectedSort, "The sort functions should sort correctly")
  }
}

func TestEmpty(t *testing.T) {
  emptyMap := map[string]int{}
  wordCounts := BuildWordCountSlice(emptyMap)
  assert.Nil(t, wordCounts, "The wordCount slice should be nil")
}

func TestSameCounts(t *testing.T) {
  expectedSort := []WordCount{{"cool",4}, {"sup",4}, {"yo",4}}
  sameMap := map[string]int{"yo":4, "cool":4, "sup":4}
  wordCounts := BuildWordCountSlice(sameMap)
  if assert.NotNil(t, wordCounts) {
    assert.Equal(t, len(wordCounts), 3, "The WordCount slice should have correct size")
    sort.Sort(ByCount(wordCounts))
    assert.Equal(t, wordCounts, expectedSort, "The sort function should break tiebreakers alphabetically")
  }
}

func TestNonAlphaWords(t *testing.T) {
  expectedSort := []WordCount{{"1",4}, {"3",4}, {"2",1}, {"4",1}}
  nonAlphaMap := map[string]int{"1":4, "2":1, "3":4, "4":1}  
  wordCounts := BuildWordCountSlice(nonAlphaMap)
  if assert.NotNil(t, wordCounts) {
    assert.Equal(t, len(wordCounts), 4, "The WordCount slice should have correct size")
    sort.Sort(ByCount(wordCounts))
    assert.Equal(t, wordCounts, expectedSort, "The sort function should sort number keys lowest first to highest")
  }
} 

func TestAlphaNumericWords(t *testing.T) {
  expectedSort := []WordCount{{"2",2}, {"b",2}, {"1",1}, {"a",1}}
  alphaNumericMap := map[string]int{"a":1, "b":2, "1":1, "2":2}
  wordCounts := BuildWordCountSlice(alphaNumericMap)
  if assert.NotNil(t, wordCounts) {
    assert.Equal(t, len(wordCounts), 4, "The WordCount slice should have correct size")
    sort.Sort(ByCount(wordCounts))
    assert.Equal(t, wordCounts, expectedSort, "The sort function should assign higher priority to numbers")
  }
}
