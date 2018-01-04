package apiserver

import (
	"bufio"
	"net"
	"strings"

	"github.com/KaiserGald/rpgApp/services/logger"
)

var conns []net.Conn
var coms chan string
var log *logger.Logger

func init() {
	conns = make([]net.Conn, 0, 10)
	coms = make(chan string)
}

// Start starts the API Server
func Start(lg *logger.Logger) error {
	log = lg
	log.Info.Log("Launching API Server")

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Error.Log("Error creating listener: %v", err)
		return err
	}
	go runServer(l)

	return nil
}

// runServer listes and accepts incoming connections, and then handles them
func runServer(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Error.Log("Error accepting connection: %v\n", err)
		}

		conns = append(conns, conn)
		go handleConnection(conn, len(conns)-1)
	}
}

// runServer listens and responds on the specified socket
func handleConnection(c net.Conn, id int) error {
	for {

		defer func() {
			log.Info.Log("Closing connection #%d\n", id)
			c.Close()
			conns[id] = nil
		}()

		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Error.Log("Error in creating Reader: %v\n")
			return err
		}

		com := string(message)
		com = strings.Trim(com, "\n")
		log.Debug.Log("Received message: %v\n", com)

		coms <- com
	}
}

// GetCommand returns the command sent by the client
func GetCommand() string {
	c := <-coms

	return c
}
