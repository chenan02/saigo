package corpus

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func testHelper(t *testing.T, expected []WordCount, m map[string]int) {
	wc := BuildWordCountSlice(m)
	if assert.NotNil(t, wc) {
		assert.Equal(t, len(wc), len(m), "BuildWordCountSlice should return a slice of correct size")
		sort.Sort(ByCount(wc))
		assert.Equal(t, wc, expected, "The sort function should sort correctly")
	}
}

func testStandard(t *testing.T) {
	expected := []WordCount{{"hello", 3}, {"hey", 3}, {"dood", 2}, {"yo", 1}}
	m := map[string]int{"hey": 3, "yo": 1, "dood": 2, "hello": 3}
	testHelper(t, expected, m)
}

func testEmpty(t *testing.T) {
	m := map[string]int{}
	wc := BuildWordCountSlice(m)
	assert.Nil(t, wc, "The wordCount slice should be nil")
}

func testSameCounts(t *testing.T) {
	expected := []WordCount{{"cool", 4}, {"sup", 4}, {"yo", 4}}
	m := map[string]int{"yo": 4, "cool": 4, "sup": 4}
	testHelper(t, expected, m)
}

func testNonAlphaWords(t *testing.T) {
	expected := []WordCount{{"1", 4}, {"3", 4}, {"2", 1}, {"4", 1}}
	m := map[string]int{"1": 4, "2": 1, "3": 4, "4": 1}
	testHelper(t, expected, m)
}

func testAlphaNumericWords(t *testing.T) {
	expected := []WordCount{{"2", 2}, {"b", 2}, {"1", 1}, {"a", 1}}
	m := map[string]int{"a": 1, "b": 2, "1": 1, "2": 2}
	testHelper(t, expected, m)
}
