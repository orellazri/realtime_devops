package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"127.0.0.1:29092"},
		Topic:   "pipeline",
		GroupID: "reader-1",
	})

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
	// conn, err := kafka.DialLeader(context.Background(), "tcp", "127.0.0.1:29092", "pipeline", 0)
	// if err != nil {
	// 	log.Fatal("failed to dial leader: ", err)
	// }

	// conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	// batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	// b := make([]byte, 10e3) // 10KB max per message
	// for {
	// 	n, err := batch.Read(b)
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Println(string(b[:n]))
	// }

	// if err := batch.Close(); err != nil {
	// 	log.Fatal("failed to close batch:", err)
	// }

	// if err := conn.Close(); err != nil {
	// 	log.Fatal("failed to close connection:", err)
	// }
}
