package testutils

import (
	"time"

	"github.com/go-feature-flag/app-api/dao/dbmodel"
)

// Bool returns a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}

// Time returns a pointer to the Time value passed in.
func Time(t time.Time) *time.Time {
	return &t
}

// Float64 returns a pointer to the float64 value passed in.
func Float64(t float64) *float64 {
	return &t
}

// Int returns a pointer to the float64 value passed in.
func Int(t int) *int {
	return &t
}

func Interface(v interface{}) *interface{} {
	return &v
}

func String(v string) *string {
	return &v
}

func JSONB(v dbmodel.JSONB) *dbmodel.JSONB {
	return &v
}
