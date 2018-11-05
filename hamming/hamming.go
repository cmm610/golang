package hamming

import (
	"strings"
	"errors"
)

func Distance(a, b string) (int, error) {
	sta := strings.Split(a, "")
	stb := strings.Split(b, "")
	l := 0

	if len(sta) == len(stb) {
		l = 0
		for i, c := range sta {
			if c != stb[i] {
				l ++
			}
		}
		return l, nil
	} else {
		return l, errors.New("New error")
	}


}
