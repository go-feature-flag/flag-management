package util

import "time"

type Clock interface {
	Now() time.Time
}

type DefaultClock struct{}

func (DefaultClock) Now() time.Time { return time.Now() }
