package main

import (
	"fmt"
	"github.com/saigo/exercise-002-tools/scanner"
	"github.com/saigo/exercise-002-tools/words"
	"os"
	"sort"
)

func main() {
	filename := os.Args[1]
	wordMap, err := scanner.ReadMap(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	wordCounts := corpus.BuildWordCountSlice(wordMap)
	sort.Sort(corpus.ByCount(wordCounts))
	// custom wordCount printLn method in WordCount library, can make Word and Count "private"
	for _, x := range wordCounts {
		fmt.Println(x.Word, x.Count)
	}
}
