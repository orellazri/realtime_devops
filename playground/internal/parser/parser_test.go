package parser

import (
	"os"
	"testing"

	"github.com/orellazri/realtime_devops/playground/internal/communicator"
	"github.com/stretchr/testify/assert"
)

var EXAMPLE_CONFIG = `
communicators:
  - sender:
      type: kafka
      address: 127.0.0.1:29092
      topic: playground-0
      delay: 0
    receiver:
      type: kafka
      address: 127.0.0.1:29092
      topic: playground-0
      delay: 0
`

func TestParseConfig(t *testing.T) {
	assert := assert.New(t)

	configPath, err := os.CreateTemp("", "")
	assert.Nil(err)

	_, err = configPath.Write([]byte(EXAMPLE_CONFIG))
	assert.Nil(err)

	cfg, err := ParseConfig(configPath.Name())
	assert.Nil(err)

	assert.Equal(len(cfg.Communicators), 1)
	assert.Equal(cfg.Communicators[0].Sender.Type, communicator.TypeKafka)
	assert.Equal(cfg.Communicators[0].Sender.Address, "127.0.0.1:29092")
	assert.Equal(cfg.Communicators[0].Sender.Topic, "playground-0")
	assert.Equal(cfg.Communicators[0].Sender.Delay, 0)

	assert.Equal(cfg.Communicators[0].Receiver.Type, communicator.TypeKafka)
	assert.Equal(cfg.Communicators[0].Receiver.Address, "127.0.0.1:29092")
	assert.Equal(cfg.Communicators[0].Receiver.Topic, "playground-0")
	assert.Equal(cfg.Communicators[0].Sender.Delay, 0)
}
