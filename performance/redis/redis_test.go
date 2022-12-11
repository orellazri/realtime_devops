package redis

import (
	"testing"
	"time"
)

func Setup(t *testing.T) *Connection {
	t.Helper()

	return NewConnection()
}

func TestWriteBenchmark(t *testing.T) {
	conn := Setup(t)

	var iterationsMap = map[int]time.Duration{
		100:  time.Duration(100 * time.Millisecond),
		1000: time.Duration(1000 * time.Millisecond),
	}

	for numIters, expectedTime := range iterationsMap {
		totalTime, err := conn.BenchmarkWrite(numIters)
		if err != nil {
			t.Fatal(err)
		}

		if totalTime > expectedTime {
			t.Errorf("total write time is too long. expected %v got %v", expectedTime, totalTime)
		}
	}
}
