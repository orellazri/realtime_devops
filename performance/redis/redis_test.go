package redis

import (
	"strconv"
	"testing"
	"time"
)

func Setup(t *testing.T) *Database {
	t.Helper()

	return NewDatabase()
}

func TestAverageWritePerformance(t *testing.T) {
	db := Setup(t)

	var totalTime time.Duration
	const numIterations int = 100

	for i := 0; i < numIterations; i++ {
		start := time.Now()
		err := db.db.Set(ctx, strconv.FormatInt(int64(i), 10), "value", 0).Err()
		if err != nil {
			t.Fatal(err)
		}
		end := time.Now()
		totalTime += end.Sub(start)
	}

	actual := totalTime / time.Duration(numIterations)
	expected := 1 * time.Millisecond

	if actual > time.Duration(expected) {
		t.Errorf("write average is too long. expected %v got %v", expected, actual)
	}
}

func TestAverageWritePerformanceAfterWarmup(t *testing.T) {
	db := Setup(t)

	var totalTime time.Duration
	numIterations := 100

	for i := 0; i < 5; i++ {
		err := db.db.Set(ctx, strconv.FormatInt(int64(i), 10), "value", 0).Err()
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < numIterations; i++ {
		start := time.Now()
		err := db.db.Set(ctx, strconv.FormatInt(int64(i), 10), "value", 0).Err()
		if err != nil {
			t.Fatal(err)
		}
		end := time.Now()
		totalTime += end.Sub(start)
	}

	actual := totalTime / time.Duration(numIterations)
	expected := 1 * time.Millisecond

	if actual > time.Duration(expected) {
		t.Errorf("write average is too long. expected %v got %v", expected, actual)
	}
}
