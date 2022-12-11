package kafka

import (
	"testing"
	"time"
)

func Setup(t *testing.T) *Connection {
	t.Helper()

	conn, err := NewConnection()
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

func Cleanup(conn *Connection, t *testing.T) {
	t.Helper()

	conn.Close()
}

func TestWriteBenchmark(t *testing.T) {
	conn := Setup(t)

	var iterationsMap = map[int]time.Duration{
		100:  time.Duration(2 * time.Second),
		1000: time.Duration(10 * time.Second),
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

	Cleanup(conn, t)
}
