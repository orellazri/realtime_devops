package redis

import (
	"context"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

type Database struct {
	DB *redis.Client
}

func NewDatabase() *Database {
	var db Database

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:55000",
		Password: "redispw",
		DB:       0,
	})
	db.DB = rdb

	return &db
}
