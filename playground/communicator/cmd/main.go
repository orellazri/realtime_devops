package main

import (
	"log"
	"sync"

	"github.com/orellazri/realtime_devops/playground/communicator/internal/communicator"
)

func main() {
	id := 0

	log.Println("Starting Kafka communicator")
	firstComm, err := communicator.NewCommunicator(
		id,
		communicator.CommunicatorDetails{
			Type:    communicator.TypeKafka,
			Address: "127.0.0.1:29092",
		},
		communicator.CommunicatorDetails{
			Type:    communicator.TypeKafka,
			Address: "127.0.0.1:29092",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		sendMessage := "hello 1"
		log.Printf("[Communicator %v] Sending: %v", firstComm.ID, sendMessage)
		err = firstComm.Send(sendMessage)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[Communicator %v] Sent: %v", firstComm.ID, sendMessage)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		receiveMessage, err := firstComm.Receive()
		log.Printf("[Communicator %v] Receiving", firstComm.ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[Communicator %v] Received: %v", firstComm.ID, receiveMessage)
	}()

	wg.Wait()
}
