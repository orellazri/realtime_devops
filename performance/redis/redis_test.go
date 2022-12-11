package redis

import (
	"testing"
	"time"
)

func Setup(t *testing.T) *Database {
	t.Helper()

	return NewDatabase()
}

func TestWriteBenchmark(t *testing.T) {
	var iterationsMap = map[int]time.Duration{
		100:  time.Duration(100 * time.Millisecond),
		1000: time.Duration(1000 * time.Millisecond),
	}

	for numIters, expectedTime := range iterationsMap {
		totalTime, err := BenchmarkWrite(numIters, false)
		if err != nil {
			t.Fatal(err)
		}

		if totalTime > expectedTime {
			t.Errorf("total write time is too long. expected %v got %v", expectedTime, totalTime)
		}
	}
}
