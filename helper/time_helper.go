package helper

import "time"

func InTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(time.Now())
}
