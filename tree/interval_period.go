package tree

import (
	"time"
)

type IntervalPeriod struct {
	Start time.Time
	End   time.Time
}

// GetStart returns the start date and time of the interval.
func (ip *IntervalPeriod) GetStart() time.Time {
	return ip.Start
}

// SetStart sets the start date and time of the interval.
func (ip *IntervalPeriod) SetStart(start time.Time) {
	ip.Start = start
}

// GetEnd returns the end date and time of the interval.
func (ip *IntervalPeriod) GetEnd() time.Time {
	return ip.End
}

// SetEnd sets the end date and time of the interval.
func (ip *IntervalPeriod) SetEnd(end time.Time) {
	ip.End = end
}
