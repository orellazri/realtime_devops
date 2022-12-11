package helpers

import "time"

type Benchmarkable interface {
	BenchmarkWrite(int) (time.Duration, error)
}
