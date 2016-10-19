package main

import (
	"fmt"
	"github.com/papaq/sortArrayOfWords/levenshteinDistance"
)

func main() {
	fmt.Println("HI!")
	lDistance.Printl()

	word1, word2 := "kitten", "cat"
	distance := lDistance.Recursive(word1, word2)
	fmt.Println("Recursive distance between", word1, "and", word2, "is:", distance)

	distance = lDistance.Iterative(word1, word2)
	fmt.Println("Iterative distance between", word1, "and", word2, "is:", distance)
	
}
