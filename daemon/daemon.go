package daemon

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KaiserGald/rpgApp/services/com/comhandler"
	"github.com/KaiserGald/rpgApp/services/com/comserver"
	"github.com/KaiserGald/rpgApp/services/logger"
	"github.com/KaiserGald/rpgApp/ui"
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

	ui.Start(l, log)
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

	log.Debug.Log("Got signal: %v, exiting.", s)
	time.Sleep(2 * time.Second)
}
