package utils

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetMessageTime(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	message := Message{ID: uuid.New(), Sent: now, Received: now.Add(30 * time.Millisecond)}
	assert.Equal(GetMessageTime(&message), time.Duration(30*time.Millisecond))
}
