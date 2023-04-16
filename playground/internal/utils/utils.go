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

func GetMessageTime(message *Message) time.Duration {
	return message.Received.Sub(message.Sent)
}


type NoMessageError struct {}
func (e *NoMessageError) Error() string {
	return "no message available yet"
}
func (e *NoMessageError) Is(target error) bool {
        return e.Error() == target.Error()
}
