package isogram

import "strings"

func IsIsogram(s string) bool {
	s = strings.Replace(s, ",", "", 80)
	s = strings.Replace(s, "-", "", 80)
	s = strings.Replace(s, " ", "", 80)
	s = strings.ToLower(s)

	strArr := strings.Split(s, "")
	strMap := map[string]int{}
	for _,c := range strArr {
		_, ok := strMap[c]
		if ok {
			return false
		} else {
			strMap[c] = 1
		}
	}
	return true
}