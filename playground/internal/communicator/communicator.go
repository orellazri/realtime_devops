package communicator

import (
	"github.com/orellazri/realtime_devops/playground/internal/clients"
)

type CommunicatorType string

const (
	TypeKafka CommunicatorType = "kafka"
)

type CommunicatorDetails struct {
	Type    CommunicatorType
	Address string
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
	senderClient, err := clients.NewKafkaClient(sender.Address)
	if err != nil {
		return nil, err
	}
	sender.client = senderClient

	receiverClient, err := clients.NewKafkaClient(receiver.Address)
	if err != nil {
		return nil, err
	}
	receiver.client = receiverClient

	return &Communicator{ID: id, Sender: sender, Receiver: receiver}, nil
}

func (communicator *Communicator) Send(message string) error {
	return communicator.Sender.client.Send(message)
}

func (communicator *Communicator) Receive() (string, error) {
	return communicator.Receiver.client.Receive()
}

func (communicator *Communicator) Close() error {
	err := communicator.Sender.client.Close()
	if err != nil {
		return err
	}
	err = communicator.Receiver.client.Close()
	if err != nil {
		return err
	}

	return nil
}
