package communicator

import (
	"github.com/orellazri/realtime_devops/playground/internal/clients"
	"github.com/orellazri/realtime_devops/playground/internal/utils"
)

const (
	TypeKafka    string = "kafka"
	TypeRedis    string = "redis"
	TypeRabbitMQ string = "rabbitmq"
	TypeEMQX     string = "emqx"
)

type CommunicatorDetails struct {
	Type    string
	Address string
	Topic   string
	Delay   int
	client  CommunicatorClient
}

type CommunicatorClient interface {
	Send(message *utils.Message) error
	Receive() (utils.Message, error)
	Close() error
}

type Communicator struct {
	ID       int
	Sender   CommunicatorDetails
	Receiver CommunicatorDetails
}

func NewCommunicator(id int, sender, receiver CommunicatorDetails) (*Communicator, error) {
	var err error
	var senderClient CommunicatorClient
	var receiverClient CommunicatorClient

	// Create sender
	if sender.Type != "" {
		switch sender.Type {
		case TypeKafka:
			senderClient, err = clients.NewKafkaClient(sender.Address, sender.Topic)
		case TypeRedis:
			senderClient, err = clients.NewRedisClient(sender.Address, sender.Topic)
		case TypeRabbitMQ:
			senderClient, err = clients.NewRabbitMQClient(sender.Address, sender.Topic)
		case TypeEMQX:
			senderClient, err = clients.NewEMQXClient(sender.Address, sender.Topic)
		}
		if err != nil {
			return nil, err
		}
		sender.client = senderClient
	}

	// Create receiver
	if receiver.Type != "" {
		switch receiver.Type {
		case TypeKafka:
			receiverClient, err = clients.NewKafkaClient(receiver.Address, receiver.Topic)
		case TypeRedis:
			receiverClient, err = clients.NewRedisClient(receiver.Address, receiver.Topic)
		case TypeRabbitMQ:
			receiverClient, err = clients.NewRabbitMQClient(receiver.Address, receiver.Topic)
		case TypeEMQX:
			receiverClient, err = clients.NewEMQXClient(receiver.Address, receiver.Topic)
		}
		if err != nil {
			return nil, err
		}
		receiver.client = receiverClient
	}

	return &Communicator{ID: id, Sender: sender, Receiver: receiver}, nil
}

func (communicator *Communicator) Send(message *utils.Message) error {
	return communicator.Sender.client.Send(message)
}

func (communicator *Communicator) Receive() (utils.Message, error) {
	return communicator.Receiver.client.Receive()
}

func (communicator *Communicator) Close() error {
	if communicator.Sender.Type != "" {
		err := communicator.Sender.client.Close()
		if err != nil {
			return err
		}
	}

	if communicator.Receiver.Type != "" {
		err := communicator.Receiver.client.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
