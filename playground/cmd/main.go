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

	for i, comm := range cfg.Communicators {
		log.Printf("Starting communicator of type %v => %v", comm.Send.Type, comm.Receive.Type)
		comm, err := communicator.NewCommunicator(i, comm.Send, comm.Receive)
		if err != nil {
			log.Fatal(err)
		}
		comms = append(comms, comm)
	}

	var wg sync.WaitGroup

	for _, comm := range comms {
		wg.Add(1)
		go func(comm *communicator.Communicator) {
			defer wg.Done()

			sendMessage := time.Now().String()
			log.Printf("[Communicator %v] Sending: %v", comm.ID, sendMessage)
			err = comm.Send(sendMessage)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("[Communicator %v] Sent: %v", comm.ID, sendMessage)
		}(comm)

		wg.Add(1)
		go func(comm *communicator.Communicator) {
			defer wg.Done()
			log.Printf("[Communicator %v] Receiving", comm.ID)
			for {
				receiveMessage, err := comm.Receive()
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("[Communicator %v] Received: %v", comm.ID, receiveMessage)
			}
		}(comm)
	}

	wg.Wait()

	for _, comm := range comms {
		log.Printf("Closing service %v", comm.ID)
		err = comm.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
