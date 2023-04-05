package communicator

import (
	"github.com/orellazri/realtime_devops/playground/internal/clients"
)

const (
	TypeKafka string = "kafka"
	TypeRedis string = "redis"
)

type CommunicatorDetails struct {
	Type    string
	Address string
	Topic   string
	Delay   int
	client  CommunicatorClient
}

type CommunicatorClient interface {
	Send(message string) error
	Receive() (string, error)
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
		}
		if err != nil {
			return nil, err
		}
		receiver.client = receiverClient
	}

	return &Communicator{ID: id, Sender: sender, Receiver: receiver}, nil
}

func (communicator *Communicator) Send(message string) error {
	return communicator.Sender.client.Send(message)
}

func (communicator *Communicator) Receive() (string, error) {
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
