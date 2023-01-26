package utils

import (
	"time"
)

func TimeRoundUp(t time.Time, round time.Duration) time.Time {
	r := t.Round(round)
	if r.Sub(t) < 0 {
		return r.Add(round)
	}
	return r
}
