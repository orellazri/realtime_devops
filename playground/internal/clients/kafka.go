package clients

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	writer *kafka.Writer
	reader *kafka.Reader
	ctx    context.Context
	cancel context.CancelFunc
}

func NewKafkaClient(address, topic string) (*KafkaClient, error) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(address),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{address},
		Topic:       topic,
		StartOffset: kafka.LastOffset,
		MinBytes:    1,
		MaxBytes:    10e6,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return &KafkaClient{writer, reader, ctx, cancel}, nil
}

func (client *KafkaClient) Send(message string) error {
	return client.writer.WriteMessages(
		client.ctx,
		kafka.Message{Key: []byte("messageKey"), Value: []byte(message)},
	)
}

func (client *KafkaClient) Receive() (string, error) {
	message, err := client.reader.ReadMessage(context.Background())
	if err != nil {
		return "", err
	}

	return string(message.Value), nil
}

func (client *KafkaClient) Close() error {
	client.cancel()

	err := client.writer.Close()
	if err != nil {
		return err
	}

	err = client.reader.Close()
	if err != nil {
		return err
	}

	return nil
}
