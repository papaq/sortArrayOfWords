package sorting

import (
	"math/rand"
	"os"
	"sort"
	"github.com/papaq/sortArrayOfWords/levenshteinDistance"
)

type WordInfo struct {
	Word string
	Distance int
}

type wdsInfoSlice []WordInfo

var (
	sliceToSort wdsInfoSlice
)

const (
	WRDMAXLEN = 10
)

func (slc wdsInfoSlice) Len() int {

	return len(slc)
}

func (slc wdsInfoSlice) Less(i, j int) bool {

	return slc[i].Distance < slc[j].Distance
}

func (slc wdsInfoSlice) Swap(i, j int) {

	slc[i], slc[j] = slc[j], slc[i]
}

func wordRand() string {	

	wrdLen := rand.Intn(2345) % (WRDMAXLEN-1) + 1
	randWord := make([]byte, wrdLen)

	for i := range randWord {
		randWord[i] = byte(rand.Intn(2345) % (25) + 98)
	}

	return string(randWord)
}

func FillRand(wordsNum int) {

	rand.Seed(int64(os.Getpid()))
	sliceToSort = make([]WordInfo, wordsNum)

	for iter := range sliceToSort {
		newWord := wordRand()
		sliceToSort[iter] = WordInfo {Word: newWord}
	}
}

func GetWords() []WordInfo {

	return []WordInfo(sliceToSort)
}

func SetWords(words []string) {

	sliceToSort = make([]WordInfo, len(words))

	for iter := range sliceToSort {
		sliceToSort[iter] = WordInfo {Word: words[iter]}
	}
}

func Sort(defaultWord string) {

	for iter := range sliceToSort {
		sliceToSort[iter].Distance = lDistance.Iterative(defaultWord, sliceToSort[iter].Word)
		// lDistance.Recursive(defaultWord, sliceToSort[iter].Word)
	}

	sort.Sort(sliceToSort)
}








