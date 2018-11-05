package luhn

import (
	"strings"
	"math"
	"strconv"
)

func checkLength(sa []string) bool {
	if len(sa) <= 1 {
		return false
	}
	return true
}

func Valid(s string) bool {
	s = strings.Replace(s, " ", "", 40)
	sa := strings.Split(s, "")
	if !checkLength(sa) {
		return false
	}

	c := 1
	var parsedNumber int
	var err error
	sum := 0
	for i := len(sa)-1; i >- 1; i-- {
		parsedNumber, err = strconv.Atoi(sa[i])
		if err != nil {
			sum += 0
			break
		}

		if math.Mod(float64(c), 2) == 0 {
			double := parsedNumber * 2
			if double > 9 {
				double = double - 9
			}

			sum += double
		} else {
			sum += parsedNumber
		}
		c ++
	}
	return math.Mod(float64(sum), 10) == 0
}