package bob

import (
	"strings"
	"unicode"
)

type Remark struct {
	remark string
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

func newRemark(remark string) Remark {
	return Remark{strings.TrimSpace(remark)}
}

func (r Remark) isShouting() bool {
	hasLetters := strings.IndexFunc(r.remark, unicode.IsLetter) >= 0
	isUpperCase := strings.ToUpper(r.remark) == r.remark

	return isUpperCase && hasLetters
}

// build Bob's response off of a remark
func Hey(remark string) string {
	r := newRemark(remark)

	switch true {
	case r.isQuestion() && r.isShouting():
		return "Calm down, I know what I'm doing!"
	case r.isQuestion() :
		return "Sure."
	case r.isShouting() :
		return "Whoa, chill out!"
	case r.remark == "":
		return "Fine. Be that way!"
	default :
		return "Whatever."
	}
}
