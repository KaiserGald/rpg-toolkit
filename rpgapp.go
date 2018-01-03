package main

import (
	"os"

	"github.com/KaiserGald/rpgApp/daemon"
	"github.com/KaiserGald/rpgApp/services/logger"
)

func processFlags() *daemon.Config {
	cfg := &daemon.Config{}
	cfg.ListenSpec = ":8080"
	return cfg
}

func main() {
	cfg := processFlags()
	log := &logger.Logger{}

	log.Init(os.Stdout, os.Stdout, os.Stderr)
	if err := daemon.Run(cfg, log); err != nil {
		log.Error.Printf("Error in main(): %v\n", err)
	}
}
