package day4

import (
	"math"
	"unicode/utf8"
)

func CountXMAS(grid []string) int {
	// Strategy:
	// Go through each row, and stop on each 'A'
	// Check the top left and bottom right have an 'M' and an 'S'
	// Check the top right and bottom left have an 'M' and an 'S'
	count := 0
	for i := range grid {
		// First and last columns no point checking for 'A's
		if i == 0 || i == len(grid)-1 {
			continue
		}
		for j := range grid[i] {
			// No point checking first or last rows for an 'A'
			if j == 0 || j == len(grid[i])-1 {
				continue
			}
			if rune(grid[i][j]) == 'A' {
				topLeft := rune(grid[i+1][j-1])
				topRight := rune(grid[i+1][j+1])
				bottomLeft := rune(grid[i-1][j-1])
				bottomRight := rune(grid[i-1][j+1])
				lrDiagonal := (topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M')
				rlDiagonal := (topRight == 'M' && bottomLeft == 'S') || (topRight == 'S' && bottomLeft == 'M')
				if lrDiagonal && rlDiagonal {
					count++
				}
			}
		}
	}
	return count
}

func CountWords(grid []string, word string) int {
	count := 0

	searchStrings := []string{}
	searchStrings = append(searchStrings, grid...)
	searchStrings = append(searchStrings, Columns(grid)...)
	searchStrings = append(searchStrings, LRDiagonals(grid)...)
	searchStrings = append(searchStrings, RLDiagonals(grid)...)

	for _, searchString := range searchStrings {
		count += CountIn(searchString, word)
		count += CountIn(reverse(searchString), word)
	}

	return count
}

func RLDiagonals(grid []string) []string {
	reversedGrid := make([]string, len(grid))
	for i := range reversedGrid {
		reversedGrid[i] = reverse(grid[i])
	}
	return LRDiagonals(reversedGrid)
}

func LRDiagonals(grid []string) []string {
	diagonals := make([]string, len(grid)+len(grid[0])-1)
	for i := range grid[0] {
		diagonal := make([]byte, int(math.Min(float64(len(grid[0])-i), float64(len(grid)))))
		for j := range diagonal {
			diagonal[j] = grid[j][i+j]
		}
		diagonals[i] = string(diagonal)
	}
	for i := range grid {
		// Skip the duplicate top left diagonal while going column wise
		if i == 0 {
			continue
		}
		diagonal := make([]byte, int(math.Min(float64(len(grid)-i), float64(len(grid[0])))))
		for j := range diagonal {
			diagonal[j] = grid[i+j][j]
		}
		diagonals[i+len(grid[0])-1] = string(diagonal)
	}
	return diagonals
}

func Columns(grid []string) []string {
	columns := make([]string, len(grid[0]))
	for j := range columns {
		column := make([]byte, len(grid))
		for i := range grid {
			column[i] = grid[i][j]
		}
		columns[j] = string(column)
	}
	return columns
}

// Counts number of occurences in a string forwards
func CountIn(s string, word string) int {
	count := 0
	lastIndex := len(s) - len(word) + 1
	if lastIndex < 0 {
		return 0
	}
	for i := range s[:lastIndex] {
		test := s[i : i+len(word)]
		if test == word {
			count++
		}
	}
	return count
}

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
