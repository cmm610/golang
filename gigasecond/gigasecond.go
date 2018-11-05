package gigasecond

import (
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	gs := time.Duration(1000000000 * time.Second)
	t = t.Add(gs)
	return t
}


//func main() {
//	AddGigasecond(time.Now())
//}