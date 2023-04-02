package clients

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	conn   *kafka.Conn
	reader *kafka.Reader
}

func NewKafkaClient(address string) (*KafkaClient, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, "playground", 0)
	if err != nil {
		return nil, err
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{address},
		Topic:       "playground",
		StartOffset: -1,
	})

	return &KafkaClient{conn, reader}, nil
}

func (client *KafkaClient) Send(message string) error {
	client.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := client.conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)
	if err != nil {
		return err
	}
	return nil
}

func (client *KafkaClient) Receive() (string, error) {
	message, err := client.reader.ReadMessage(context.Background())
	if err != nil {
		return "", err
	}

	return string(message.Value), nil
}
