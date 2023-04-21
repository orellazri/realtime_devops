package main

import (
	"os"

	"github.com/orellazri/realtime_devops/playground/internal/controller"
	"github.com/orellazri/realtime_devops/playground/internal/parser"

	log "github.com/sirupsen/logrus"
)

var logLevels = map[string]log.Level{
	"":        log.TraceLevel,
	"trace":   log.TraceLevel,
	"warning": log.WarnLevel,
	"error":   log.ErrorLevel,
}

func main() {
	logLevel := os.Getenv("LOG_LEVEL")
	log.SetLevel(logLevels[logLevel])

	cfg, err := parser.ParseConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	controller.HandleCommunicators(cfg)
}
