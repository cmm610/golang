package reverse

import (
	"bytes"
)

func String(s string) string {
	var output bytes.Buffer

	runes := []rune(s)
	for i := range runes {
		output.WriteRune(runes[len(runes) - 1 - i])
	}

	return output.String()
}