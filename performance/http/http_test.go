package http

import (
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	numIterations := 100

	totalTime, err := Benchmark(numIterations)
	if err != nil {
		t.Fatal(err)
	}

	if totalTime/time.Duration(numIterations) < time.Duration(1*time.Microsecond) {
		t.Errorf("write average is too long. expected %v got %v", 1*time.Microsecond, totalTime/time.Duration(numIterations))
	}
}
