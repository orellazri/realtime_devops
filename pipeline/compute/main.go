package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Get environemnt varaibles
	kafkaUrl := os.Getenv("KAFKA_URL")
	if kafkaUrl == "" {
		log.Fatal("KAFKA_URL env var not found")
	}

	rabbitmqUrl := os.Getenv("RABBITMQ_URL")
	if rabbitmqUrl == "" {
		log.Fatal("RABBITMQ_URL env var not found")
	}

	// Initialize Kafka
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{kafkaUrl},
		Topic:       "pipeline",
		StartOffset: -1,
	})

	// Initialize RabbitMQ
	conn, err := amqp.Dial(rabbitmqUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"compute", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Receive Kafka messages
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		start := time.Now()
		msg := strings.Split(string(m.Value), ",")
		msgTime := msg[0]
		x, err := strconv.Atoi(msg[1])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(msg[2])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[Kafka - %v] Received (offset %d) = %s,%d,%d\n", time.Since(start), m.Offset, msgTime, x, y)

		// Perform some computation
		start = time.Now()
		x *= 10
		y *= 10
		log.Printf("\t[Core - %v] Computed %d,%d", time.Since(start), x, y)

		// Send message to RabbitMQ
		start = time.Now()
		body := fmt.Sprintf("%s,%d,%d", msgTime, x, y)
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\t\t[RabbitMQ - %v] Sent: %s\n", time.Since(start), body)
	}
}
