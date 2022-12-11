package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

type Database struct {
	db *redis.Client
}

func NewDatabase() *Database {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:55000",
		Password: "redispw",
		DB:       0,
	})

	db := &Database{db: rdb}
	return db
}

func BenchmarkWrite(numIterations int) (time.Duration, error) {
	db := NewDatabase()

	var totalTime time.Duration
	for i := 0; i < numIterations; i++ {
		start := time.Now()
		err := db.db.Set(ctx, strconv.FormatInt(int64(i), 10), "value", 0).Err()
		if err != nil {
			return time.Duration(0), err
		}
		end := time.Now()
		totalTime += end.Sub(start)
	}

	return totalTime, nil
}
