package parser

import (
	"os"

	"github.com/orellazri/realtime_devops/playground/internal/communicator"
	"gopkg.in/yaml.v3"
)

type service struct {
	Send    communicator.CommunicatorDetails
	Receive communicator.CommunicatorDetails
}

type config struct {
	Communicators []service
}

func ParseConfig() (*config, error) {
	cfg := config{}

	data, err := os.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
