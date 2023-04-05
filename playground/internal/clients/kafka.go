package clients

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	writer *kafka.Writer
	reader *kafka.Reader
}

func NewKafkaClient(address, topic string) (*KafkaClient, error) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(address),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
		BatchSize:              1,
		RequiredAcks:           kafka.RequireOne,
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{address},
		Topic:       topic,
		StartOffset: kafka.LastOffset,
		MinBytes:    1,
		MaxBytes:    10e6,
	})

	return &KafkaClient{writer, reader}, nil
}

func (client *KafkaClient) Send(message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return client.writer.WriteMessages(
		ctx,
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
