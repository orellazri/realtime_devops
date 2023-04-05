package clients

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
	topic  string
}

func NewRedisClient(address, topic string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	return &RedisClient{client, topic}, nil
}

func (client *RedisClient) Send(message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return client.client.Set(ctx, client.topic, message, 0).Err()
}

func (client *RedisClient) Receive() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message, err := client.client.Get(ctx, client.topic).Result()
	if err != nil {
		return "", err
	}

	return string(message), nil
}

func (client *RedisClient) Close() error {
	return nil
}
