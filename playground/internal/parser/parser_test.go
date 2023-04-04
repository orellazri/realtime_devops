package parser

import (
	"os"
	"testing"

	"github.com/orellazri/realtime_devops/playground/internal/communicator"
	"github.com/stretchr/testify/assert"
)

var EXAMPLE_CONFIG = `
communicators:
  - send:
      type: kafka
      address: 127.0.0.1:29092
    receive:
      type: kafka
      address: 127.0.0.1:29092
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
	assert.Equal(cfg.Communicators[0].Send.Type, communicator.TypeKafka)
	assert.Equal(cfg.Communicators[0].Send.Address, "127.0.0.1:29092")
	assert.Equal(cfg.Communicators[0].Receive.Type, communicator.TypeKafka)
	assert.Equal(cfg.Communicators[0].Receive.Address, "127.0.0.1:29092")
}
