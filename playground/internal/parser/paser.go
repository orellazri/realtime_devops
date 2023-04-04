package parser

import (
	"os"

	"github.com/orellazri/realtime_devops/playground/internal/communicator"
	"gopkg.in/yaml.v3"
)

type comm struct {
	Sender   communicator.CommunicatorDetails
	Receiver communicator.CommunicatorDetails
}

type config struct {
	Communicators []comm
}

func ParseConfig(path string) (*config, error) {
	cfg := config{}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
