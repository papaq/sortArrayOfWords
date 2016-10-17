package lDistance

import (
	"fmt";
)

type TwoWords struct {
	word1, word2	[]byte
	w1length, w2length byte
}


var (
	words TwoWords
	matrix []byte
)

func Printl() {
	fmt.Println("aaah?")
}

func min2(a, b byte) byte {
	if a < b {
		return a
	} 
	return b
}

func min3(a, b, c byte) byte {
	return min2(min2(a, b), c)
}

func xyToxByRows(currCol, currRow, nCols byte) byte {
	return nCols*(currRow)+currCol
}

func xyToxByColumns(currCol, currRow, nRows byte) byte {
	return nRows*(currCol)+currRow
}

func Recursive(word1, word2 string) byte {
	words = TwoWords{[]byte(word1), []byte(word2), byte(len(word1)), byte(len(word2))}
	
	matrix = make([]byte, (words.w1length+1)*(words.w2length+1))	

	if 0 > 1 {
		return matrix[0]
	}
	
	return runRecursive(words.w2length, words.w1length)
}

func runRecursive(x, y byte) byte {
	
	if cell := matrix[xyToxByRows(x, y, words.w2length+1)]; cell != 0 {
		return cell
	}

	if x*y == 0 {
		matrix[xyToxByRows(x, y, words.w2length+1)] = x+y
		return x+y
	}

	var condition byte
	if words.word1[y-1] == words.word2[x-1] {
		condition = 0
	} else {
		condition = 1
	}

	cell := min3(runRecursive(x-1, y)+1,
			runRecursive(x, y-1)+1,
			runRecursive(x-1, y-1) + condition,
			)
	
	matrix[xyToxByRows(x, y, words.w2length+1)] = cell	
	
	return cell
} 








