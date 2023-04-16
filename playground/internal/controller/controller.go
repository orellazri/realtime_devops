package controller

import (
	"log"
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

	// Start communicators
	log.Println("🚀 Starting communicators")
	for i := 0; i < 5; i++ {
		startPipeline(comms)
	}

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

func startPipeline(comms []*communicator.Communicator) {
	for _, comm := range comms {
		var msg utils.Message
		var err error
		if comm.Receiver.Type == "" && comm.Sender.Type != "" {
			// First communicator - should generate a message and send it
			msg = utils.Message{ID: uuid.New(), Sent: time.Now()}
			err = comm.Send(msg)
			if err != nil {
				log.Fatalf("[%v] Error while sending: %v", comm.ID, err);
			}
			log.Printf("[%v] (%v) ➡️ %v", comm.ID, comm.Sender.Topic, msg.ID)
		} else if comm.Receiver.Type != "" && comm.Sender.Type != "" {
			// Middle communicators - should receive a message and send it
			msg, err = comm.Receive()
			if err != nil {
				log.Fatalf("[%v] Error while receiving: %v", comm.ID, err);
			}
			log.Printf("[%v] (%v) ⬅️ %v", comm.ID, comm.Receiver.Topic, msg.ID)

			err = comm.Send(msg)
			if err != nil {
				log.Fatalf("[%v] Error while sending: %v", comm.ID, err);
			}
			log.Printf("[%v] (%v) ➡️ %v", comm.ID, comm.Sender.Topic, msg.ID)
		} else if comm.Receiver.Type != "" && comm.Sender.Type == "" {
			// Last communicator - should only receive a message
			msg, err = comm.Receive()
			if err != nil {
				log.Fatalf("[%v] Error while receiving: %v", comm.ID, err);
			}
			log.Printf("[%v] (%v) ⬅️ %v", comm.ID, comm.Receiver.Topic, msg.ID)
			msg.Received = time.Now()
			messages = append(messages, msg)
		} else {
			log.Fatalf("Communicator %v is invalid", comm.ID);
		}
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
