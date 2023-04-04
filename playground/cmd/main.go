package main

import (
	"log"
	"sync"
	"time"

	"github.com/orellazri/realtime_devops/playground/internal/communicator"
	"github.com/orellazri/realtime_devops/playground/internal/parser"
)

func main() {
	cfg, err := parser.ParseConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Create communicators
	log.Println("🆕 Creating communicators...")
	var comms []*communicator.Communicator
	for i, comm := range cfg.Communicators {
		log.Printf("Creating communicator of type %v (%v) => %v (%v)", comm.Sender.Type, comm.Sender.Topic, comm.Receiver.Type, comm.Receiver.Topic)
		comm, err := communicator.NewCommunicator(i, comm.Sender, comm.Receiver)
		if err != nil {
			log.Fatal(err)
		}
		comms = append(comms, comm)
	}

	// Start communicators
	log.Println("🚀 Starting communicators...")
	var wg sync.WaitGroup
	for _, comm := range comms {
		wg.Add(1)
		go func(comm *communicator.Communicator) {
			defer wg.Done()

			for {
				sendMessage := time.Now().String()
				log.Printf("[%v] (%v): Sending %v", comm.ID, comm.Sender.Topic, sendMessage)
				err = comm.Send(sendMessage)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("[%v] (%v) ➡️ %v", comm.ID, comm.Sender.Topic, sendMessage)
				time.Sleep(time.Duration(comm.Sender.Delay) * time.Second)
			}
		}(comm)

		wg.Add(1)
		go func(comm *communicator.Communicator) {
			defer wg.Done()
			log.Printf("[%v] (%v): Receiving", comm.ID, comm.Receiver.Topic)
			for {
				receiveMessage, err := comm.Receive()
				if err != nil {
					log.Printf("[%v] Error while receiving: %v", comm.ID, err)
					break
				}
				log.Printf("[%v] (%v) ⬅️ %v", comm.ID, comm.Receiver.Topic, receiveMessage)
				time.Sleep(time.Duration(comm.Receiver.Delay) * time.Second)
			}
		}(comm)
	}

	// Wait for communicators to finish
	wg.Wait()

	// Close communicators
	log.Println("🚪 Closing communicators...")
	for _, comm := range comms {
		log.Printf("Closing communicator %v", comm.ID)
		err = comm.Close()
		if err != nil {
			log.Printf("Error while closing communicator %v: %v", comm.ID, err)
		}
	}
}
