package clients

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	conn *kafka.Conn
}

func NewKafkaClient(address string) (*KafkaClient, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, "playground", 0)
	if err != nil {
		return nil, err
	}

	return &KafkaClient{conn}, nil
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
	return "", nil
}
