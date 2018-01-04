package daemon

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/KaiserGald/rpgApp/services/apiserver"
	"github.com/KaiserGald/rpgApp/services/logger"
	"github.com/KaiserGald/rpgApp/ui"
)

// Config contains configuration information for the server
type Config struct {
	ListenSpec string
	DevMode    bool
}

// Run starts up the server daemon
func Run(cfg *Config, log *logger.Logger) error {

	log.Info.Log("Starting HTTP on %s\n", cfg.ListenSpec)
	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Error.Log("Error creating listener: %v\n", err)
		return err
	}

	ui.Start(l, log)
	apiserver.Start(log)
	com := apiserver.GetCommand()
	if com == "stop" {
		log.Notice.Log("Stop command received, shutting down...\n")
		os.Exit(0)
	}

	waitForSignal(log)

	return nil
}

func waitForSignal(log *logger.Logger) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch

	log.Info.Log("Got signal: %v, exiting.", s)
}
