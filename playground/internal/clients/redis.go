package clients

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
	topic  string
	ctx    context.Context
	cancel context.CancelFunc
}

func NewRedisClient(address, topic string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return &RedisClient{client, topic, ctx, cancel}, nil
}

func (client *RedisClient) Send(message string) error {
	return client.client.Set(client.ctx, client.topic, message, 0).Err()
}

func (client *RedisClient) Receive() (string, error) {
	message, err := client.client.Get(client.ctx, client.topic).Result()
	if err != nil {
		return "", err
	}

	return string(message), nil
}

func (client *RedisClient) Close() error {
	client.cancel()
	return nil
}
