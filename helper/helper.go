package helper

import "time"

func Now() time.Time {
	now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return now
}
