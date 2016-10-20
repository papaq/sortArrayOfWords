package main

import (
	"github.com/papaq/sortArrayOfWords/inAndOut"
	"github.com/papaq/sortArrayOfWords/sorting"
)

func main() {

	defword := "slovakia"
	file := "countries.txt"
	wordsToSort := iao.ReadFile(file)
	sorting.SetWords(wordsToSort)

	iao.LineToConsole("Sorted list of words according to Levenshtein distance to \"" +
		defword + "\":")

	sorting.Sort(defword)

	sortedWordsInfo := sorting.GetWords()

	for _, wordInfo := range sortedWordsInfo {
		iao.WordInfoToConsole(wordInfo.Word, wordInfo.Distance)
	}
}
