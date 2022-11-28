package http

import (
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	var iterationsMap = map[int]time.Duration{
		100:  time.Duration(20 * time.Millisecond),
		1000: time.Duration(70 * time.Millisecond),
	}

	for numIters, expectedTime := range iterationsMap {
		totalTime, err := Benchmark(numIters)
		if err != nil {
			t.Fatal(err)
		}

		if totalTime > expectedTime {
			t.Errorf("total write time is too long. expected %v got %v", expectedTime, totalTime)
		}
	}

}
