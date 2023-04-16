package clients

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/orellazri/realtime_devops/playground/internal/utils"
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

func (client *RedisClient) Send(message *utils.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return client.client.RPush(ctx, client.topic, data).Err()
}

func (client *RedisClient) Receive() (utils.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message, err := client.client.RPop(ctx, client.topic).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return utils.Message{}, &utils.NoMessageError{}
		}
		return utils.Message{}, err
	}

	var data utils.Message
	err = json.Unmarshal([]byte(message), &data)
	if err != nil {
		return utils.Message{}, err
	}

	return data, nil
}

func (client *RedisClient) Close() error {
	return nil
}
