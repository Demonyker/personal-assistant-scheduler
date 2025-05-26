package main

import (
	"log"

	"github.com/Demonyker/personal-assistant-scheduler/config"
	"github.com/Demonyker/personal-assistant-scheduler/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
