package utils

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID   uuid.UUID
	Sent time.Time
}
