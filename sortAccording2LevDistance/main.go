package main

import (
	"fmt"
	"github.com/papaq/sortArrayOfWords/levenshteinDistance"
)

func main() {
	fmt.Println("HI!")
	lDistance.Printl()

	word1, word2 := "kitten", "sitting"
	distance := lDistance.Recursive(word1, word2)
	fmt.Println("distance between", word1, "and", word2, "is", distance)

}
