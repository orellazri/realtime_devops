package utils

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID
	Sent     time.Time
	Received time.Time
}

func GetMessageTime(message Message) time.Duration {
	return message.Received.Sub(message.Sent)
}
