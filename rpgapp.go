package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/KaiserGald/rpgApp/daemon"
	"github.com/KaiserGald/rpgApp/services/logger"
)

var dev bool
var port int
var verbose bool
var quiet bool
var log *logger.Logger
var logLevel int

func processCLI() *daemon.Config {
	cfg := &daemon.Config{}

	processFlags(cfg)

	configureDaemon(cfg)

	return cfg
}

func processFlags(cfg *daemon.Config) {
	flag.BoolVar(&dev, "dev", false, "sets server to dev mode")
	flag.BoolVar(&dev, "d", false, "sets server to dev mode")

	flag.IntVar(&port, "port", 8080, "sets listen port for server")
	flag.IntVar(&port, "p", 8080, "sets listen port for server")

	flag.BoolVar(&verbose, "verbose", false, "sets server log to Verbose mode")
	flag.BoolVar(&verbose, "v", false, "sets server log to Verbose mode")

	flag.BoolVar(&quiet, "quiet", false, "sets server log to Verbose mode")
	flag.BoolVar(&quiet, "q", false, "sets server log to Verbose mode")

	flag.Parse()

}

func configureDaemon(cfg *daemon.Config) {
	if dev {
		cfg.DevMode = true
		cfg.ListenSpec = ":" + strconv.Itoa(port)
		logLevel = 0
	} else {
		cfg.DevMode = false
		cfg.ListenSpec = ":" + os.Getenv("PORT")
		logLevel = 2
	}

	if verbose {
		logLevel = 1
	}

	if quiet {
		logLevel = 3
	}
}

func main() {
	log := &logger.Logger{}
	cfg := processCLI()

	log.Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr, logLevel)

	if err := daemon.Run(cfg, log); err != nil {
		log.Error.Log("Error in main(): %v\n", err)
	}
}
