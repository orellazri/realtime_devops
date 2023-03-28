package emqx

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Connection struct {
	client *mqtt.Client
}

func NewConnection() *Connection {
	mqtt.ERROR = log.New(os.Stderr, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetClientID("emqx_test_client")

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(messageCallback)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to a topic
	if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	return &Connection{client: &c}
}

func (conn *Connection) Close() error {
	time.Sleep(1 * time.Second)

	// Unsubscribe
	token := (*conn.client).Unsubscribe("testtopic/#")
	token.Wait()
	err := token.Error()
	if err != nil {
		return err
	}

	// Disconnect
	(*conn.client).Disconnect(250)
	time.Sleep(1 * time.Second)
	return nil
}

var messageCallback mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	// Dummy message handler so the messages will be acknowledged
}

func (conn *Connection) BenchmarkWrite(numIterations int) (time.Duration, error) {
	var totalTime time.Duration
	for i := 0; i < numIterations; i++ {
		start := time.Now()
		token := (*conn.client).Publish("testtopic/1", 0, false, strconv.FormatInt(int64(i), 10))
		_ = token.Wait()
		if token.Error() != nil {
			log.Fatal(token.Error())
		}
		end := time.Now()
		totalTime += end.Sub(start)
	}

	return totalTime, nil
}
