// Package daemon
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package daemon

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KaiserGald/logger"
	"github.com/KaiserGald/unlichtServer/services/com/comhandler"
	"github.com/KaiserGald/unlichtServer/services/com/comserver"
	"github.com/KaiserGald/unlichtServer/ui"
)

var log *logger.Logger

// Config contains configuration information for the server
type Config struct {
	ListenSpec string
	DevMode    bool
}

// Run starts up the server daemon
func Run(cfg *Config, lg *logger.Logger) error {
	log = lg

	log.Notice.Log("Starting HTTP on %s\n", cfg.ListenSpec)
	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Error.Log("Error creating listener: %v\n", err)
		return err
	}

	err = ui.Start(l, log)
	if err != nil {
		log.Error.Log("Error Starting UI.\n")
		return err
	}
	comserver.Start(log)
	comhandler.Start(log)

	log.Notice.Log("Server up and running.\n")

	waitForSignal()

	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch

	log.Debug.Log("Got signal: %v, exiting...", s)
	time.Sleep(2 * time.Second)
}
