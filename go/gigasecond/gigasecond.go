package gigasecond

import (
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)
	return t
}
