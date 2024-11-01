package clock

import (
	"time"
)

type Clock interface {
	Now() time.Duration
}

type Timer struct {
	startTime time.Time
}

func NewTimer() Clock {
	return &Timer{
		startTime: time.Now(),
	}
}

func (r *Timer) Now() time.Duration {
	return time.Since(r.startTime)
}
