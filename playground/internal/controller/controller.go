package controller

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

func HandleCommunicators(cfg *parser.Config) {
	// Create communicators
	log.Println("🆕 Creating communicators...")
	var comms []*communicator.Communicator
	for i, comm := range cfg.Communicators {
		comms = append(comms, createCommunicator(i, comm))
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start communicators
	log.Println("🚀 Starting communicators...")
	for _, comm := range comms {
		startCommunicator(ctx, &wg, comm)
	}

	// SIGTERM handler
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Received SIGTERM. Stopping...")
		cancel()
	}()

	wg.Wait()

	log.Println("🚪 Closing communicators...")
	for _, comm := range comms {
		closeCommunicator(comm)
	}
}

func createCommunicator(index int, parserComm parser.ConfigCommunicator) *communicator.Communicator {
	log.Printf("Creating communicator of type %v (%v) => %v (%v)", parserComm.Sender.Type, parserComm.Sender.Topic, parserComm.Receiver.Type, parserComm.Receiver.Topic)
	comm, err := communicator.NewCommunicator(index, parserComm.Sender, parserComm.Receiver)
	if err != nil {
		log.Fatal(err)
	}
	return comm
}

func startCommunicator(ctx context.Context, wg *sync.WaitGroup, comm *communicator.Communicator) {
	wg.Add(2)

	go func(comm *communicator.Communicator) {
		defer wg.Done()
		log.Printf("[%v] (%v): Sending", comm.ID, comm.Sender.Topic)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				sendMessage := time.Now().Format(time.RFC3339Nano)
				err := comm.Send(sendMessage)
				if err != nil {
					log.Printf("[%v] Error while sending: %v", comm.ID, err)
					break
				}
				log.Printf("[%v] (%v) ➡️ %v", comm.ID, comm.Sender.Topic, sendMessage)
				time.Sleep(time.Duration(comm.Sender.Delay) * time.Millisecond)
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
				receiveTime, err := time.Parse(time.RFC3339Nano, receiveMessage)
				if err != nil {
					log.Printf("[%v] Error while parsing timestamp: %v", comm.ID, err)
					break
				}
				log.Printf("[%v] (%v) ⬅️ %v", comm.ID, comm.Receiver.Topic, receiveMessage)
				log.Printf("    %v", time.Since(receiveTime))
				time.Sleep(time.Duration(comm.Receiver.Delay) * time.Millisecond)
			}
		}
	}(comm)
}

func closeCommunicator(comm *communicator.Communicator) {
	log.Printf("Closing communicator %v", comm.ID)
	err := comm.Close()
	if err != nil {
		log.Printf("Error while closing communicator %v: %v", comm.ID, err)
	}
}
