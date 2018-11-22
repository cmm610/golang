//Package raindrops takes an integer value and "translates" it to a particular raindrop sound by determining the integer's prime factorization.

package raindrops

import (
	"strconv"
)

const testVersion = 2

var translations = []struct {
	Num  int
	Word string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(amount int) string {
	s := ""

	for _, trans := range translations {
		if amount%trans.Num == 0 {
			s += trans.Word
		}
	}

	if s == "" {
		s = strconv.Itoa(amount)
	}

	return s
}