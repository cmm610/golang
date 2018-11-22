package scrabble

import (
	"strings"
)
var wordVals = map[string]int {
	"A": 1,
	"E": 1,
	"I": 1,
	"O": 1,
	"U": 1,
	"L": 1,
	"N": 1,
	"R": 1,
	"S": 1,
	"T": 1,
	"D": 2,
	"G": 2,
	"B": 3,
	"C": 3,
	"M": 3,
	"P": 3,
	"F": 4,
	"H": 4,
	"V": 4,
	"W": 4,
	"Y": 4,
	"K": 5,
	"J": 8,
	"X": 8,
	"Z": 10,
	"Q": 10,
}
func Score(w string) int {
	wa := strings.Split(w, "")
	wordMap := map[string]int{}
	for _, c := range wa {
		_, ok := wordMap[c]
		if ok {
			wordMap[c] += 1
		} else {
			wordMap[c] = 1
		}
	}

	total := 0
	for char, val := range wordMap {
		total += wordVals[strings.ToUpper(char)] * val
	}
	return total
}