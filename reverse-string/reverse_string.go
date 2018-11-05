package reverse

import (
	"strings"
)

func String(s string) string {
	ss := strings.Split(s, "")
	l := len(ss)
	mod := int(l/2)
	retArr := make([]string, l)
	key := 0
	for i, runeVal := range s {

	}
	for i := 0; i<mod; i++ {
		key = l - 1 - i
		retArr[i] = ss[key]
		retArr[key] = ss[i]
		if i == mod - 1 {
			retArr[i+1] = ss[i+1]
		}
	}
	return strings.Join(retArr, "")
}