package comhandler

import (
	"net"
	"os"
	"time"

	"github.com/KaiserGald/logger"
	"github.com/KaiserGald/unlichtServer/services/com/comserver"
)

const (
	stop    string = "stop"
	restart string = "restart"
	online  string = "online"
)

var (
	com      string
	conn     *net.TCPConn
	log      *logger.Logger
	p        *os.Process
	shutdown bool
)

// Start start the Command Handler
func Start(lg *logger.Logger) {
	log = lg
	shutdown = false
	var err error
	p, err = os.FindProcess(os.Getpid())
	if err != nil {
		log.Error.Log("Error finding pid: %v\n", err)
	}
	go handle()
}

// handle handles any command coming into the command server
func handle() {
	for {
		com, conn := comserver.GetCommand()
		switch com {
		case stop:
			log.Notice.Log("Stop command received, shutting server down...\n")
			comserver.Respond(conn, "stop\n")
			shutdown = true
			err := p.Signal(os.Interrupt)
			if err != nil {
				log.Error.Log("Error emitting interrupt signal: %v\n", err)
			}
		case restart:
			log.Notice.Log("Restart command received, restarting server now...\n")
			comserver.Respond(conn, "restart\n")
			comserver.Kill()
			shutdown = true
			time.Sleep(5 * time.Second)
			err := p.Signal(os.Interrupt)
			if err != nil {
				log.Error.Log("Error emitting interrupt signal: %v\n", err)
			}
		case online:
			if !shutdown {
				log.Debug.Log("Online status request received. Responding to request with status 'online'.")
				comserver.Respond(conn, "online\n")
			} else {
				log.Debug.Log("Online status request received. Responding to request with status 'offline'.")
				comserver.Respond(conn, "offline\n")
			}

		default:
			log.Debug.Log("Unknown Command Received.\n")
			comserver.Respond(conn, "unknown\n")
		}
	}
}
