package util_test

import (
	"testing"
	"time"

	"github.com/go-feature-flag/app-api/util"
)

func TestDefaultClock_Now(t *testing.T) {
	var clock util.Clock = util.DefaultClock{}
	now := clock.Now()

	// Allow a small margin of error for the time difference
	if time.Since(now) > time.Second {
		t.Errorf("Expected time close to now, but got %v", now)
	}
}
