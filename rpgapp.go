package main

import (
	"flag"
	"os"

	"github.com/KaiserGald/rpgApp/daemon"
	"github.com/KaiserGald/rpgApp/services/logger"
)

var dev bool
var port int
var verbose bool

func processCLI() *daemon.Config {
	cfg := &daemon.Config{}

	processFlags(cfg)

	return cfg
}

func processFlags(cfg *daemon.Config) {
	flag.BoolVar(&dev, "dev", false, "sets server to dev mode")
	flag.BoolVar(&dev, "d", false, "sets server to dev mode")

	flag.IntVar(&port, "port", 8080, "sets listen port for server")
	flag.IntVar(&port, "p", 8080, "sets listen port for server")

	flag.BoolVar(&verbose, "verbose", false, "sets server log to Verbose mode")
	flag.BoolVar(&verbose, "v", false, "sets server log to Verbose mode")

	flag.Parse()

	if dev {
		cfg.DevMode = true

	}
}

func main() {
	cfg := processCLI()
	log := &logger.Logger{}

	log.Init(os.Stdout, os.Stdout, os.Stderr)
	if err := daemon.Run(cfg, log); err != nil {
		log.Error.Printf("Error in main(): %v\n", err)
	}
}
