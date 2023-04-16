package clients

import (
	"context"
	"encoding/json"
	"time"

	"github.com/orellazri/realtime_devops/playground/internal/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    *amqp.Queue
}

func NewRabbitMQClient(address, topic string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Fetch one message at a time
	err = ch.Qos(1, 0, false)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		topic, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQClient{conn, ch, &q}, nil
}

func (client *RabbitMQClient) Send(message *utils.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return client.ch.PublishWithContext(ctx,
		"",            // exchange
		client.q.Name, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		},
	)
}

func (client *RabbitMQClient) Receive() (utils.Message, error) {
	msgs, err := client.ch.Consume(
		client.q.Name, // queue
		"",            // consumer
		false,         // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		return utils.Message{}, err
	}

	message := <-msgs

	var data utils.Message
	err = json.Unmarshal(message.Body, &data)
	if err != nil {
		return utils.Message{}, err
	}

	return data, nil
}

func (client *RabbitMQClient) Close() error {
	err := client.conn.Close()
	if err != nil {
		return err
	}

	err = client.ch.Close()
	if err != nil {
		return err
	}

	return nil
}
