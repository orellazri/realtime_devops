package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

type Connection struct {
	db *redis.Client
}

func NewConnection() *Connection {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:55000",
		Password: "redispw",
		DB:       0,
	})

	return &Connection{db: rdb}
}

func (conn *Connection) BenchmarkWrite(numIterations int) (time.Duration, error) {
	var totalTime time.Duration
	for i := 0; i < numIterations; i++ {
		start := time.Now()
		err := conn.db.Set(ctx, strconv.FormatInt(int64(i), 10), "value", 0).Err()
		if err != nil {
			return time.Duration(0), err
		}
		end := time.Now()
		totalTime += end.Sub(start)
	}

	return totalTime, nil
}
