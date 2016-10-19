package lDistance

import (
	"testing";
)

func slicesEqualByte(a, b []byte) bool {

	if len(a) != len(b) {
		return false
	}

	for i, cell := range a {
		if cell != b[i] {
		return false
		}
	}

	return true
}

func slicesEqualInt(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	for i, cell := range a {
		if cell != b[i] {
		return false
		}
	}

	return true
}

func TestslicesEqualInt(t *testing.T) {

	type twoSlices struct {
		slice1, slice2 []int;
	}

	causal := []struct {
		in twoSlices;
		out bool;
	} {
		{twoSlices{[]int{}, []int{}}, true},
		{twoSlices{[]int{}, []int{12, 24}}, false},
		{twoSlices{[]int{12}, []int{12, 24}}, false},
		{twoSlices{[]int{12, 2004}, []int{12, 2004}}, true},
	}

	for _, set := range causal {
		got := slicesEqualInt(set.in.slice1, set.in.slice2)

		if set.out != got {
			t.Errorf("slicesEqualInt(%q, %q) == %q, want %q", 
			set.in.slice1, set.in.slice2, got, set.out)
		}
	}
}

func TestslicesEqualByte(t *testing.T) {

	type twoSlices struct {
		slice1, slice2 []byte;
	}

	causal := []struct {
		in twoSlices;
		out bool;
	} {
		{twoSlices{[]byte{}, []byte{}}, true},
		{twoSlices{[]byte{}, []byte{12, 24}}, false},
		{twoSlices{[]byte{12}, []byte{12, 24}}, false},
		{twoSlices{[]byte{12, 24}, []byte{12, 24}}, true},
	}

	for _, set := range causal {
		got := slicesEqualByte(set.in.slice1, set.in.slice2)

		if set.out != got {
			t.Errorf("slicesEqualByte(%q, %q) == %q, want %q",
			set.in.slice1, set.in.slice2, got, set.out)
		}
	}
}

func TestbuildTwoWords(t *testing.T) {

	type inWords struct {
		word1, word2 string;
	}

	causal := []struct {
		in inWords;
		out TwoWords;
	} {
		{inWords{"", ""}, 
			TwoWords{word1: []byte{}, word2: []byte{}}},
		{inWords{"word", "word1"}, 
			TwoWords{
				[]byte{119, 111, 114, 100}, 
				[]byte{119, 111, 114, 100, 49}, 5, 5,
		}},
		{inWords{"x", "y"}, TwoWords{[]byte{120}, []byte{121}, 1, 1}},
	}

	for _, set := range causal {
		got := buildTwoWords(set.in.word1, set.in.word2)

		if slicesEqualByte(set.out.word1, got.word1) ||
			slicesEqualByte(set.out.word2, got.word2) ||
			set.out.w1length != got.w1length ||
			set.out.w2length != got.w2length {
			t.Errorf("buildTwoWords(%q, %q) == %q, want %q", 
			set.in.word1, set.in.word2, got, set.out)
		}
	}
}

func TestbuildMatrix(t *testing.T) {

	type inWords struct {
		word1, word2 string;
	}

	causal := []struct {
		in inWords;
		out []int;
	} {
		{inWords{"", ""}, []int{0}},
		{inWords{"w", "o"}, []int{0, 0, 0, 0}},
	}

	for _, set := range causal {
		got := buildMatrix(set.in.word1, set.in.word2)

		if slicesEqualInt(set.out, got) {
			t.Errorf("buildMatrix(%q, %q) == %q, want %q", 
			set.in.word1, set.in.word2, got, set.out)
		}
	}
}

func Testmin2(t *testing.T) {
	
	type values struct {
		value1, value2 int;
	}

	causal := []struct {
		in values;
		out int;
	} {
		{values{5, 12}, 5},
		{values{123, 123}, 123},
	}

	for _, set := range causal {
		got := min2(set.in.value1, set.in.value2)

		if set.out != got {
			t.Errorf("min2(%q, %q) == %q, want %q", 
			set.in.value1, set.in.value2, got, set.out)
		}
	}
}

func Testmin3(t *testing.T) {	

	type values struct {
		value1, value2, value3 int;
	}

	causal := []struct {
		in values;
		out int;
	} {
		{values{5, 12, 13}, 5},
		{values{123, 123, 123}, 123},
		{values{123, 0, -273}, -273},
	}

	for _, set := range causal {
		got := min3(set.in.value1, set.in.value2, set.in.value3)

		if set.out != got {
			t.Errorf("min3(%q, %q, %q) == %q, want %q", 
			set.in.value1, set.in.value2, set.in.value3,
			got, set.out)
		}
	}
}

func TestxyToxByRows(t *testing.T) {

	type inParams struct {
		currCol, currRow, nCols int;
	}

	causal := []struct {
		in inParams;
		out int;
	} {
		{inParams{0, 0, 13}, 0},
		{inParams{1, 5, 5}, 10},
		{inParams{0, 20, 3}, 60},
	}

	for _, set := range causal {
		got := xyToxByRows(set.in.currCol, set.in.currRow, set.in.nCols)

		if set.out != got {
			t.Errorf("xyToxByRows(%q, %q, %q) == %q, want %q", 
			set.in.currCol, set.in.currRow, set.in.nCols, 
			got, set.out)
		}
	}
}

func TestxyToxByColumns(t *testing.T) {

	type inParams struct {
		currCol, currRow, nRows int;
	}

	causal := []struct {
		in inParams;
		out int;
	} {
		{inParams{0, 0, 13}, 0},
		{inParams{5, 1, 5}, 10},
		{inParams{20, 0, 3}, 60},
	}

	for _, set := range causal {
		got := xyToxByColumns(set.in.currCol, set.in.currRow, set.in.nRows)

		if set.out != got {
			t.Errorf("xyToxByColumns(%q, %q, %q) == %q, want %q", 
			set.in.currCol, set.in.currRow, set.in.nRows, 
			got, set.out)
		}
	}
}

func TestboolToInt(t *testing.T) {

	causal := []struct {
		in bool;
		out int;
	} {
		{true, 1},
		{false, 0},
	}

	for _, set := range causal {
		got := boolToInt(set.in)

		if set.out != got {
			t.Errorf("boolToInt(%q) == %q, want %q", 
			set.in, got, set.out)
		}
	}
}

func TestRecursive(t *testing.T) {

	type inWords struct {
		word1, word2 string;
	}

	causal := []struct {
		in inWords;
		out int;
	} {
		{inWords{"", ""}, 0},
		{inWords{"w", "o"}, 1},
		{inWords{"o", "o"}, 0},
		{inWords{"cat", "kitkat"}, 4},
	}

	for _, set := range causal {
		got := Recursive(set.in.word1, set.in.word2)

		if set.out != got {
			t.Errorf("Recursive(%q, %q) == %q, want %q", 
			set.in.word1, set.in.word2, got, set.out)
		}
	}
}

func TestIterative(t *testing.T) {

	type inWords struct {
		word1, word2 string;
	}

	causal := []struct {
		in inWords;
		out int;
	} {
		{inWords{"", ""}, 0},
		{inWords{"w", "o"}, 1},
		{inWords{"o", "o"}, 0},
		{inWords{"cat", "kitkat"}, 4},
	}

	for _, set := range causal {
		got := Iterative(set.in.word1, set.in.word2)

		if set.out != got {
			t.Errorf("Iterative(%q, %q) == %q, want %q", 
			set.in.word1, set.in.word2, got, set.out)
		}
	}
}

func TestcountMatrixCell(t *testing.T) {

	type inWords struct {
		word1, word2 string;
	}

	type inCoordinates struct {
		x, y int;
	}

	type inValues struct {
		words inWords;
		coordinates inCoordinates;
	}

	causal := []struct {
		in inValues;
		out int;
	} {
		{inValues{
			inWords{"a", "b"}, inCoordinates{1, 1},
		}, 1},
		{inValues{
			inWords{"tom", "none"}, inCoordinates{2, 3},
		}, 2},
	}

	for _, set := range causal {
		words = buildTwoWords(set.in.words.word1, set.in.words.word2)
		matrix = buildMatrix(set.in.words.word1, set.in.words.word2)
		
		got := countMatrixCell(set.in.coordinates.x, 
			set.in.coordinates.y, len(set.in.words.word2))

		if set.out != got {
			t.Errorf("countMatrixCell(%q, %q, %q) == %q, want %q",
				set.in.coordinates.x, set.in.coordinates.y, 
				len(set.in.words.word2),
				got, set.out)
		}
	}
}

