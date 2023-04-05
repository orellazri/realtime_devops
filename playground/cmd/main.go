package main

import (
	"log"

	"github.com/orellazri/realtime_devops/playground/internal/controller"
	"github.com/orellazri/realtime_devops/playground/internal/parser"
)

func main() {
	cfg, err := parser.ParseConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	controller.HandleCommunicators(cfg)
}
