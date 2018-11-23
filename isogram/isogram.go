// Package isogram contains the solution to the corresponding EXERCISM problem.
package isogram

import "unicode"

// IsIsogram returns true iff s is an isogram.
func IsIsogram(s string) bool {
	seen := make(map[rune]bool)
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		_, ok := seen[unicode.ToUpper(r)]
		if ok {
			return false
		}
		seen[unicode.ToUpper(r)] = true
	}
	return true
}