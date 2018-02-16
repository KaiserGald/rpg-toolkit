// Package main
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/KaiserGald/logger"
	"github.com/KaiserGald/unlicht-server/daemon"
)

var dev bool
var port int
var verbose bool
var quiet bool
var color bool
var log *logger.Logger
var logLevel int

func processCLI() *daemon.Config {
	log = logger.New()
	cfg := &daemon.Config{}

	processFlags(cfg)

	configureDaemon(cfg)

	return cfg
}

func processFlags(cfg *daemon.Config) {
	flag.BoolVar(&dev, "dev", false, "sets server to dev mode")
	flag.BoolVar(&dev, "d", false, "sets server to dev mode")

	flag.BoolVar(&color, "color", false, "Colors the server output.")
	flag.BoolVar(&color, "c", false, "Colors the server output.")

	flag.IntVar(&port, "port", 8080, "sets listen port for server")
	flag.IntVar(&port, "p", 8080, "sets listen port for server")

	flag.BoolVar(&verbose, "verbose", false, "sets server log to Verbose mode")
	flag.BoolVar(&verbose, "v", false, "sets server log to Verbose mode")

	flag.BoolVar(&quiet, "quiet", false, "sets server log to Verbose mode")
	flag.BoolVar(&quiet, "q", false, "sets server log to Verbose mode")

	flag.Parse()

}

func configureDaemon(cfg *daemon.Config) {
	env := os.Getenv("RUN_ENV")
	log.SetLogLevel(logger.All)
	log.Debug.Log(env)
	log.Info.Log("Configuring server daemon...")
	if env == "DEV" {
		cfg.DevMode = true
		cfg.ListenSpec = ":" + strconv.Itoa(port)
		log.Debug.Log("Started in dev mode.")
	} else if env == "PROD" {
		cfg.DevMode = false
		cfg.ListenSpec = ":" + os.Getenv("PORT")
		log.SetLogLevel(logger.Normal)
		log.Debug.Log("Started in production mode.")
	}

	if verbose {
		log.SetLogLevel(logger.Verbose)
		log.Debug.Log("Started in verbose mode.")
	}

	if quiet {
		log.SetLogLevel(logger.ErrorsOnly)
		log.Debug.Log("Started in quiet mode.")
	}
	if color {
		log.Debug.Log("Started in colored output mode.")
	}
	log.ShowColor(color)
}

func main() {
	cfg := processCLI()
	log.Notice.Log("Starting app '%v'!", os.Getenv("BINARY_NAME"))

	log.Info.Log("Starting server daemon...")
	if err := daemon.Run(cfg, log); err != nil {
		log.Error.Log("Error in main(): %v", err)
	}
}
