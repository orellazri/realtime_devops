package controller

import (
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/orellazri/realtime_devops/playground/internal/communicator"
	"github.com/orellazri/realtime_devops/playground/internal/parser"
	"github.com/orellazri/realtime_devops/playground/internal/utils"
	log "github.com/sirupsen/logrus"
)

var messages []*utils.Message

func HandleCommunicators(cfg *parser.Config) {
	// Create communicators
	log.Println("🆕 Creating communicators")
	var comms []*communicator.Communicator
	for i, comm := range cfg.Communicators {
		comms = append(comms, createCommunicator(i, comm))
	}

	// Start pipeline
	log.Println("🚀 Starting pipeline")
	for i := 0; i < cfg.Meta.NumMessages; i++ {
		startPipeline(comms)
	}

	log.Println("🚪 Closing communicators")
	for _, comm := range comms {
		closeCommunicator(comm)
	}

	log.Println("📈 Statistics")
	printStats(cfg.Meta.NumMessages)
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
		var currentMessage *utils.Message = &utils.Message{}
		var lastMessage *utils.Message
		if len(messages) > 0 {
			lastMessage = messages[len(messages)-1]
		} else {
			lastMessage = &utils.Message{}
		}

		if comm.Receiver.Type == "" && comm.Sender.Type != "" {
			// First communicator - should generate a message and send it
			currentMessage = &utils.Message{ID: uuid.New(), Sent: time.Now()}
			sendMessage(comm, currentMessage)
			messages = append(messages, currentMessage)
		} else if comm.Receiver.Type != "" && comm.Sender.Type != "" {
			// Middle communicators - should receive a message and send it
			for currentMessage.ID != lastMessage.ID {
				log.Tracef("\t[%v] Skipped %v", comm.ID, currentMessage.ID)
				currentMessage = receiveMessage(comm)
			}

			sendMessage(comm, currentMessage)
		} else if comm.Receiver.Type != "" && comm.Sender.Type == "" {
			// Last communicator - should only receive a message
			for currentMessage.ID != lastMessage.ID {
				log.Tracef("\t[%v] Skipped %v", comm.ID, currentMessage.ID)
				currentMessage = receiveMessage(comm)
			}

			lastMessage.Received = time.Now()
			log.Printf("\t[%v] %v", comm.ID, time.Since(lastMessage.Sent))
			log.Println("---------------")
		} else {
			log.Fatalf("Communicator %v is not configured correctly", comm.ID)
		}
	}
}

func sendMessage(comm *communicator.Communicator, message *utils.Message) {
	err := comm.Send(message)
	if err != nil {
		log.Fatalf("[%v] Error while sending: %v", comm.ID, err)
	}
	log.Printf("[%v] (%v) ➡️ %v", comm.ID, comm.Sender.Topic, message.ID)
	time.Sleep(time.Duration(comm.Sender.Delay) * time.Millisecond)
}

func receiveMessage(comm *communicator.Communicator) *utils.Message {
	message, err := comm.Receive()
	if err != nil {
		log.Fatalf("[%v] Error while receiving: %v", comm.ID, err)
	}
	log.Printf("[%v] (%v) ⬅️ %v", comm.ID, comm.Receiver.Topic, message.ID)
	time.Sleep(time.Duration(comm.Receiver.Delay) * time.Millisecond)
	return &message
}

func closeCommunicator(comm *communicator.Communicator) {
	log.Printf("Closing communicator %v", comm.ID)
	err := comm.Close()
	if err != nil {
		log.Printf("Error while closing communicator %v: %v", comm.ID, err)
	}
}

func printStats(numMessages int) {
	// Create array to store all message times
	allTimes := make([]time.Duration, numMessages)
	for i, message := range messages {
		allTimes[i] = utils.GetMessageTime(message)
	}

	// Sort times array
	sort.Slice(allTimes, func(i, j int) bool { return allTimes[i] < allTimes[j] })

	// Calculate total time (sum of all times)
	var totalTime time.Duration
	for _, time := range allTimes {
		totalTime += time
	}

	log.Printf("Sent %v messages", numMessages)
	log.Printf("Total time: %v", totalTime)
	log.Printf("Fastest time: %v", allTimes[0])
	log.Printf("Slowest time: %v", allTimes[len(allTimes) - 1])
	log.Printf("Average time: %v", totalTime/time.Duration(numMessages))
}
