package bob

import (
	"regexp"
	"strings"
)

// build Bob's response off of a remark
func Hey(remark string) string {
	remark = strings.Replace(remark, " ", "", 44)
	bs := []byte (remark)
	c, _ := regexp.Match("[A-Z]{4}", bs)
	q, _ := regexp.Match("\\?$", bs)
	n, _ := regexp.Match("[1-9]", bs)
	e, _ := regexp.Match("\\!", bs)
	as, _ := regexp.Match("\\t", bs)
	if c && q {
		return "Calm down, I know what I'm doing!"
	} else if n&&q  || q {
		return "Sure."
	} else if n&&e  || c {
		return "Whoa, chill out!"
	} else if remark == "" ||  as {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
