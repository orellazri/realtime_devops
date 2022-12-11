package kafka

import (
	"context"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type Connection struct {
	conn *kafka.Conn
}

func NewConnection() (*Connection, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "192.168.1.211:9092", "benchmark", 0)
	if err != nil {
		return nil, err
	}

	return &Connection{conn: conn}, nil
}

func (conn *Connection) Close() error {
	err := conn.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (conn *Connection) BenchmarkWrite(numIterations int) (time.Duration, error) {
	var totalTime time.Duration
	for i := 0; i < numIterations; i++ {
		start := time.Now()
		conn.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		_, err := conn.conn.WriteMessages(
			kafka.Message{Value: []byte(strconv.FormatInt(int64(i), 10))},
		)
		if err != nil {
			return time.Duration(0), err
		}
		end := time.Now()
		totalTime += end.Sub(start)
	}

	return totalTime, nil
}
