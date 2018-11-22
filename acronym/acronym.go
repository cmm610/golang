package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	r := strings.Replace(s, "-", " ", -1)
	sa := strings.Fields(r)
	ret := ""
	for _, word := range sa {
		ret += strings.ToUpper(string(word[0]))
	}

	return ret
}
