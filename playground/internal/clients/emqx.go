package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/orellazri/realtime_devops/playground/internal/utils"
)

type EMQXClient struct {
	client *mqtt.Client
	topic  string
}

var messagesMutex sync.Mutex
var messages []utils.Message

func NewEMQXClient(address, topic string) (*EMQXClient, error) {
	mqtt.ERROR = log.New(os.Stderr, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(address).SetClientID(uuid.NewString())
	opts.SetKeepAlive(60 * time.Second)

	// Set the message callback handler
	opts.SetDefaultPublishHandler(messageCallback)
	opts.SetPingTimeout(1 * time.Second)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	_ = token.Wait()
	err := token.Error()
	if err != nil {
		return nil, err
	}

	// Subscribe to topic
	token = client.Subscribe(fmt.Sprintf("%s/#", topic), 0, onMessageReceived)
	_ = token.Wait()
	err = token.Error()
	if err != nil {
		return nil, err
	}

	return &EMQXClient{client: &client, topic: topic}, nil
}

func (client *EMQXClient) Send(message *utils.Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	token := (*client.client).Publish(fmt.Sprintf("%s/1", client.topic), 0, false, data)
	_ = token.Wait()
	err = token.Error()
	if err != nil {
		return err
	}

	return nil
}

func (client *EMQXClient) Receive() (utils.Message, error) {
	for len(messages) == 0 {
		time.Sleep(1 * time.Millisecond)
	}

	messagesMutex.Lock()
	message := messages[0]
	messages = messages[1:]
	messagesMutex.Unlock()

	return message, nil
}

func (client *EMQXClient) Close() error {
	time.Sleep(1 * time.Second)

	// Unsubscribe
	token := (*client.client).Unsubscribe(fmt.Sprintf("%s/#", client.topic))
	_ = token.Wait()
	err := token.Error()
	if err != nil {
		return err
	}

	// Disconnect
	(*client.client).Disconnect(250)
	time.Sleep(1 * time.Second)
	return nil
}

var messageCallback mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
	// Dummy message handler so the messages will be acknowledged
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	var data utils.Message
	err := json.Unmarshal(message.Payload(), &data)
	if err != nil {
		log.Fatal(err)
	}

	messagesMutex.Lock()
	messages = append(messages, data)
	messagesMutex.Unlock()
}
