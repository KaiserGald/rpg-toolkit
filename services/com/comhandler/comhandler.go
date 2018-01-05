package comhandler

import (
	"net"
	"os"

	"github.com/KaiserGald/rpgApp/services/com/comserver"
	"github.com/KaiserGald/rpgApp/services/logger"
)

const (
	stop string = "stop"
)

var (
	com  string
	conn *net.TCPConn
	log  *logger.Logger
	p    *os.Process
)

// Start start the Command Handler
func Start(lg *logger.Logger) {
	log = lg
	var err error
	p, err = os.FindProcess(os.Getpid())
	if err != nil {
		log.Error.Log("Error finding pid: %v\n", err)
	}
	go handle()
}

// handle handles any command coming into the api
func handle() {
	for {
		com, conn := comserver.GetCommand()
		switch com {
		case stop:
			log.Notice.Log("Stop command received, shutting server down...\n")
			comserver.Respond(conn, "stop\n")

			err := p.Signal(os.Interrupt)
			if err != nil {
				log.Error.Log("Error emitting interrupt signal: %v\n", err)
			}
		default:
			log.Debug.Log("Unknown Command Received.\n")
			comserver.Respond(conn, "unknown\n")
		}
	}
}
