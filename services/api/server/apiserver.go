package apiserver

import (
	"bufio"
	"net"
	"strings"

	"github.com/KaiserGald/rpgApp/services/logger"
)

var conns []*net.TCPConn
var coms chan string
var log *logger.Logger
var service string

func init() {
	conns = make([]*net.TCPConn, 0, 10)
	coms = make(chan string, 5)
}

// Start starts the API Server
func Start(lg *logger.Logger) error {
	log = lg
	service = ":8081"

	log.Info.Log("Launching API Server")

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		log.Error.Log("Error in resolving TCP address: %v", err)
		return err
	}

	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Error.Log("Error creating listener: %v", err)
		return err
	}
	go runServer(l)

	return nil
}

// runServer listes and accepts incoming connections, and then handles them
func runServer(l *net.TCPListener) {
	//var conn net.TCPConn
	//var err error
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Error.Log("Error accepting connection: %v\n", err)
		} else {
			conns = append(conns, conn)
			go handleConnection(conn, len(conns)-1)
		}
	}
}

// runServer listens and responds on the specified socket
func handleConnection(c *net.TCPConn, id int) error {

	defer func() {
		log.Debug.Log("Closing connection #%d.\n", id)
		c.Close()
		conns[id] = nil
	}()

	log.Info.Log("Connection %v established.", id)
	for {
		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Info.Log("Connection to client %v lost.\n", id)
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

// Respond responds to the client with a message
func Respond(c *net.TCPConn, s string) {
	c.Write([]byte(s))
}
