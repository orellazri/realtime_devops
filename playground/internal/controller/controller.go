package controller

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/orellazri/realtime_devops/playground/internal/communicator"
	"github.com/orellazri/realtime_devops/playground/internal/parser"
	"github.com/orellazri/realtime_devops/playground/internal/utils"
)

var messages []utils.Message

func HandleCommunicators(cfg *parser.Config) {
	// Create communicators
	log.Println("🆕 Creating communicators")
	var comms []*communicator.Communicator
	for i, comm := range cfg.Communicators {
		comms = append(comms, createCommunicator(i, comm))
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start communicators
	log.Println("🚀 Starting communicators")
	for _, comm := range comms {
		startCommunicator(ctx, &wg, comm)
	}

	// SIGTERM handler
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Received SIGTERM. Stopping playground")
		cancel()
	}()

	wg.Wait()

	log.Println("🚪 Closing communicators")
	for _, comm := range comms {
		closeCommunicator(comm)
	}

	log.Println("📈 Statistics")
	printStats()
}

func createCommunicator(index int, parserComm parser.ConfigCommunicator) *communicator.Communicator {
	log.Printf("Creating communicator [%v] ➡️ %v (%v) | ⬅️ %v (%v)", index, parserComm.Sender.Type, parserComm.Sender.Topic, parserComm.Receiver.Type, parserComm.Receiver.Topic)
	comm, err := communicator.NewCommunicator(index, parserComm.Sender, parserComm.Receiver)
	if err != nil {
		log.Fatal(err)
	}
	return comm
}

func startCommunicator(ctx context.Context, wg *sync.WaitGroup, comm *communicator.Communicator) {
	if comm.Sender.Type != "" {
		wg.Add(1)
		go func(comm *communicator.Communicator) {
			defer wg.Done()
			log.Printf("[%v] (%v): Sending", comm.ID, comm.Sender.Topic)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					time.Sleep(time.Duration(comm.Sender.Delay) * time.Millisecond)
					sendMessage := utils.Message{ID: uuid.New(), Sent: time.Now()}
					err := comm.Send(sendMessage)
					if err != nil {
						log.Printf("[%v] Error while sending: %v", comm.ID, err)
						break
					}
					log.Printf("[%v] (%v) ➡️ %v", comm.ID, comm.Sender.Topic, sendMessage.ID)
				}
			}
		}(comm)
	}

	if comm.Receiver.Type != "" {
		wg.Add(1)
		go func(comm *communicator.Communicator) {
			defer wg.Done()
			log.Printf("[%v] (%v): Receiving", comm.ID, comm.Receiver.Topic)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					time.Sleep(time.Duration(comm.Receiver.Delay) * time.Millisecond)
					receiveMessage, err := comm.Receive()
					if err != nil {
						// Skip the error if this is a no message error (meaning we are still waiting
						// to receive a message)
						if errors.Is(err, &utils.NoMessageError{}) {
							continue
						}

						log.Printf("[%v] Error while receiving: %v", comm.ID, err)
						break
					}
					receiveMessage.Received = time.Now()
					messages = append(messages, receiveMessage)

					log.Printf("[%v] (%v) ⬅️ %v", comm.ID, comm.Receiver.Topic, receiveMessage.ID)
					log.Printf("    %v", utils.GetMessageTime(receiveMessage))
				}
			}
		}(comm)
	}
}

func closeCommunicator(comm *communicator.Communicator) {
	log.Printf("Closing communicator %v", comm.ID)
	err := comm.Close()
	if err != nil {
		log.Printf("Error while closing communicator %v: %v", comm.ID, err)
	}
}

func printStats() {
	fastestMessage := utils.Message{ID: uuid.New(), Sent: time.Time{}, Received: time.Now()}
	var slowestMessage utils.Message
	for _, message := range messages {
		if utils.GetMessageTime(message) < utils.GetMessageTime(fastestMessage) {
			fastestMessage = message
		}
		if utils.GetMessageTime(message) > utils.GetMessageTime(slowestMessage) {
			slowestMessage = message
		}
	}

	log.Printf("Fastest message took: %v\n", utils.GetMessageTime(fastestMessage))
	log.Printf("Slowest message took: %v\n", utils.GetMessageTime(slowestMessage))
}
