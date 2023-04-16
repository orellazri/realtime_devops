package clients

import (
	"context"
	"encoding/json"
	"time"

	"github.com/orellazri/realtime_devops/playground/internal/utils"
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

func (client *KafkaClient) Send(message *utils.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return client.writer.WriteMessages(
		ctx,
		kafka.Message{Key: []byte("messageKey"), Value: []byte(data)},
	)
}

func (client *KafkaClient) Receive() (utils.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message, err := client.reader.ReadMessage(ctx)
	if err != nil {
		return utils.Message{}, err
	}

	var data utils.Message
	err = json.Unmarshal(message.Value, &data)
	if err != nil {
		return utils.Message{}, err
	}

	return data, nil
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
