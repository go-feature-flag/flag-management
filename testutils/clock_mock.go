package testutils

import "time"

type ClockMock struct {
}

func (c ClockMock) Now() time.Time {
	return time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
}
