package goutil

import (
	"time"
)

func TimeFormat(t time.Time, formatStr string) string {
	var timeString = t.Format(formatStr)
	return timeString
}
