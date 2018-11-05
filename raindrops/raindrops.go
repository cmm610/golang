package raindrops

import (
	"math"
	"strconv"
)

func Convert(i int) string {
	f := float64(i)
	ret := ""

	if math.Mod(f, 3) == 0 {
		ret += "Pling"
	}
	if math.Mod(f, 5) == 0 {
		ret += "Plang"
	}
	if math.Mod(f, 7) == 0 {
		ret += "Plong"
	}
	if ret == "" {
		ret = strconv.Itoa(i)
	}
	return ret
}