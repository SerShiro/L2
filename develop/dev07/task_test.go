package dev

import (
	"testing"
	"time"
)

func Test_or(t *testing.T) {
	start := time.Now()

	ch1 := sig(2 * time.Hour)
	ch2 := sig(5 * time.Minute)
	ch3 := sig(1 * time.Second)
	ch4 := sig(1 * time.Hour)
	ch5 := sig(1 * time.Minute)

	doneChannel := or(ch1, ch2, ch3, ch4, ch5)

	<-doneChannel

	duration := time.Since(start)
	expectedDuration := time.Minute
	if duration > expectedDuration {
		t.Errorf("Expected duration: %v, Actual duration: %v", expectedDuration, duration)
	}
}
