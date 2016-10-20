package lDistance

type TwoWords struct {
	word1, word2 []byte
	w1length, w2length int
}

var (
	words TwoWords
	matrix []int
)

func buildTwoWords(word1, word2 string) TwoWords {

	return TwoWords{[]byte(word1), []byte(word2), len(word1), len(word2)}
}

func buildMatrix(word1, word2 string) []int {

	return make([]int, (len(word1)+1)*(len(word2)+1))
}

func min2(a, b int) int {

	if a < b {
		return a
	} 
	return b
}

func min3(a, b, c int) int {

	return min2(min2(a, b), c)
}

func xyToxByRows(currCol, currRow, nCols int) int {

	return nCols*(currRow)+currCol
}

func xyToxByColumns(currCol, currRow, nRows int) int {

	return nRows*(currCol)+currRow
}

func boolToInt(b bool) int {

	if b {
		return 1
	}
	return 0
}

func Recursive(word1, word2 string) int {

	words = buildTwoWords(word1, word2)
	matrix = buildMatrix(word1, word2)

	return runRecursive(words.w2length, words.w1length)
}

func runRecursive(x, y int) int {

	columns := words.w2length + 1
	
	if cell := matrix[xyToxByRows(x, y, columns)]; cell != 0 {
		return cell
	}

	if 0 == x * y {
		matrix[xyToxByRows(x, y, columns)] = x + y
		return x + y
	}

	condition := boolToInt(words.word1[y - 1] != words.word2[x - 1])

	cell := min3(runRecursive(x - 1, y) + 1,
			runRecursive(x, y - 1) + 1,
			runRecursive(x - 1, y - 1) + condition,
			)
	
	matrix[xyToxByRows(x, y, columns)] = cell	
	
	return cell
} 

func Iterative(word1, word2 string) int {

	words = buildTwoWords(word1, word2)
	matrix = buildMatrix(word1, word2)
	columns := words.w2length + 1

	// Fill the left (N0) column with {0, 1, 2, ...}
	for y := 1; y <= words.w1length; y=y+1 {
		matrix[xyToxByRows(0, y, columns)] = y
	}
	
	// Fill the upper (N0) row with {0, 1, 2, ...}
	for x := 1; x <= words.w2length; x=x+1 {
		matrix[xyToxByRows(x, 0, columns)] = x
	}

	for x := 1; x <= words.w2length; x++ {
		for y := 1; y <= words.w1length; y++ {
			matrix[xyToxByRows(x, y, columns)] = countMatrixCell(x, y, columns)
		}
	}

	return matrix[xyToxByRows(words.w2length, words.w1length, columns)];
}

func countMatrixCell(x, y, columns int) int {

	var condition = boolToInt(words.word1[y - 1] != words.word2[x - 1])

	return min3(matrix[xyToxByRows(x - 1, y, columns)] + 1,
		matrix[xyToxByRows(x, y - 1, columns)] + 1,
		matrix[xyToxByRows(x - 1, y - 1, columns)] + condition,
		)
}

