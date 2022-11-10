package gotools

import (
	"fmt"
	"time"
)

type timeOpt struct{}

// CalcTimestampDuration 计算时间间隔，精确至日
func (opt timeOpt) CalcTimestampDuration(timeBefore, timeAfter time.Time) string {
	duration := ""
	if timeBefore.Unix() >= timeAfter.Unix() {
		return duration
	}

	yearBefore, monthBefore, dayBefore := timeBefore.Date()
	yearAfter, monthAfter, dayAfter := timeAfter.Date()

	yearDiff, monthDiff, dayDiff := yearAfter-yearBefore, monthAfter-monthBefore, dayAfter-dayBefore
	if yearDiff > 0 {
		duration += fmt.Sprintf("%d年", yearDiff)
	}
	if monthDiff > 0 {
		duration += fmt.Sprintf("%d月", monthDiff)
	}
	if dayDiff > 0 {
		duration += fmt.Sprintf("%d日", dayDiff)
	}

	return duration
}
