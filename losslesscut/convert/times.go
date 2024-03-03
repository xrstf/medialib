package convert

import (
	"strings"
	"time"
)

func durationToTime(d time.Duration) time.Time {
	return time.Time{}.Add(d)
}

func durationToTimestamp(d time.Duration) string {
	return timeToTimestamp(durationToTime(d))
}

func secondsToTime(s float64) time.Time {
	return time.UnixMicro(int64(s * 1000000))
}

func secondsToTimestamp(s float64) string {
	return timeToTimestamp(secondsToTime(s))
}

func timeToTimestamp(t time.Time) string {
	return strings.ReplaceAll(t.UTC().Format("15:04:05.000"), ".", ":")
}
