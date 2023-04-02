package communicator

import (
	"log"

	"github.com/orellazri/realtime_devops/playground/communicator/internal/clients"
)

type CommunicatorType int

const (
	TypeKafka CommunicatorType = iota
)

type CommunicatorDetails struct {
	Type    CommunicatorType
	Address string
	client  CommunicatorClient
}

type CommunicatorClient interface {
	Send(message string) error
	Receive() (string, error)
}

type Communicator struct {
	id       int
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

	return &Communicator{Sender: sender, Receiver: receiver}, nil
}

func (communicator *Communicator) Send(message string) error {
	log.Printf("[Communicator %v] Sending: %v", communicator.id, message)
	return communicator.Sender.client.Send(message)
}

func (communicator *Communicator) Receive() (string, error) {
	return "", nil
}
