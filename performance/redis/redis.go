package redis

import (
	"context"

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
