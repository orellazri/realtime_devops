package main

import (
	"log"

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
	firstComm.Send("hello 1")
}
