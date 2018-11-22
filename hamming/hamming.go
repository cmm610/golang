package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	distance := 0

	if len(a) != len(b) {
		return distance, errors.New("New error")

	}
	for i := range a {
		if a[i] != b[i] {
			distance ++
		}
	}
	return distance, nil
}
