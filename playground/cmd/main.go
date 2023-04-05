package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start communicators
	log.Println("🚀 Starting communicators...")
	for _, comm := range comms {
		wg.Add(2)

		go func(comm *communicator.Communicator) {
			defer wg.Done()
			log.Printf("[%v] (%v): Sending", comm.ID, comm.Sender.Topic)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					sendMessage := time.Now().Format(time.RFC3339)
					err = comm.Send(sendMessage)
					if err != nil {
						log.Printf("[%v] Error while sending: %v", comm.ID, err)
						break
					}
					log.Printf("[%v] (%v) ➡️ %v", comm.ID, comm.Sender.Topic, sendMessage)
					time.Sleep(time.Duration(comm.Sender.Delay) * time.Second)
				}

			}
		}(comm)

		go func(comm *communicator.Communicator) {
			defer wg.Done()
			log.Printf("[%v] (%v): Receiving", comm.ID, comm.Receiver.Topic)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					receiveMessage, err := comm.Receive()
					if err != nil {
						log.Printf("[%v] Error while receiving: %v", comm.ID, err)
						break
					}
					receiveTime, err := time.Parse(time.RFC3339, receiveMessage)
					if err != nil {
						log.Printf("[%v] Error while parsing timestamp: %v", comm.ID, err)
						break
					}
					log.Printf("[%v] (%v) ⬅️ %v", comm.ID, comm.Receiver.Topic, receiveMessage)
					log.Printf("    %v", time.Since(receiveTime))
					time.Sleep(time.Duration(comm.Receiver.Delay) * time.Second)
				}
			}
		}(comm)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Received SIGTERM. Stopping...")
		cancel()
	}()

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
