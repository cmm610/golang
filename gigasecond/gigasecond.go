package gigasecond

import (
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	gs := time.Duration(1e9 * time.Second)
	t = t.Add(gs)
	return t
}
