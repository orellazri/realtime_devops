package main

import (
	"log"
	"sync"
	"time"

	"github.com/orellazri/realtime_devops/playground/internal/communicator"
	"github.com/orellazri/realtime_devops/playground/internal/parser"
)

func main() {
	cfg, err := parser.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	var comms []*communicator.Communicator

	for i, service := range cfg.Communicators {
		log.Println("Starting communicator of type %v => %v", service.Send.Type, service.Receive.Type)
		comm, err := communicator.NewCommunicator(i, service.Send, service.Receive)
		if err != nil {
			log.Fatal(err)
		}
		comms = append(comms, comm)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		sendMessage := time.Now().String()
		log.Printf("[Communicator %v] Sending: %v", comms[0].ID, sendMessage)
		err = comms[0].Send(sendMessage)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[Communicator %v] Sent: %v", comms[0].ID, sendMessage)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("[Communicator %v] Receiving", comms[0].ID)
		for {
			receiveMessage, err := comms[0].Receive()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("[Communicator %v] Received: %v", comms[0].ID, receiveMessage)
		}
	}()

	wg.Wait()

	for _, comm := range comms {
		log.Printf("Closing service %v", comm.ID)
		err = comm.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
