package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	r := strings.Replace(s, "-", " ", 40)
	arr := strings.Split(r, " ")
	ret := ""
	for i := 0; i < len(arr); i++ {
		ret += strings.ToUpper(arr[i][0:1])
	}

	return ret
}
