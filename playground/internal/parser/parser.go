package parser

import (
	"os"

	"github.com/orellazri/realtime_devops/playground/internal/communicator"
	"gopkg.in/yaml.v3"
)

type ConfigCommunicator struct {
	Sender   communicator.CommunicatorDetails
	Receiver communicator.CommunicatorDetails
}

type Config struct {
	Communicators []ConfigCommunicator
}

func ParseConfig(path string) (*Config, error) {
	cfg := Config{}

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
