package scrabble

import "unicode"

func Score(word string) int {
	count := 0
	for i := 0; i < len(word); i++ {
		l := unicode.ToUpper(rune(word[i]))
		if l == 'A' || l == 'E' || l == 'I' || l == 'O' || l == 'U' || l == 'L' || l == 'N' || l == 'R' || l == 'S' || l == 'T' {
			count = count + 1
		} else if l == 'D' || l == 'G' {
			count = count + 2
		} else if l == 'B' || l == 'C' || l == 'M' || l == 'P' {
			count = count + 3
		} else if l == 'F' || l == 'H' || l == 'V' || l == 'W' || l == 'Y' {
			count = count + 4
		} else if l == 'K' {
			count = count + 5
		} else if l == 'J' || l == 'X' {
			count = count + 8
		} else if l == 'Q' || l == 'Z' {
			count = count + 10
		}
	}
	return count
}
