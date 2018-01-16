package comserver

import (
	"bufio"
	"net"
	"strings"

	"github.com/KaiserGald/rpgApp/services/logger"
)

var conns []*net.TCPConn
var comch chan Command
var log *logger.Logger
var service string
var kill bool

func init() {
	conns = make([]*net.TCPConn, 0, 10)
	comch = make(chan Command, 5)
}

// Command contains a command in the form of a string and a pointer to net.TCPConn
type Command struct {
	Command string
	Conn    *net.TCPConn
}

// Start starts the Command Server
func Start(lg *logger.Logger) error {
	log = lg
	service = ":8081"

	log.Info.Log("Launching Command Server")

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
func runServer(l *net.TCPListener) error {
	//var conn net.TCPConn
	//var err error
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Error.Log("Error accepting connection: %v\n", err)
			return err
		}
		conns = append(conns, conn)
		go handleConnection(conn, len(conns)-1)
		if kill {
			l.Close()
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

		m := string(message)
		m = strings.Trim(m, "\n")
		log.Debug.Log("Received message: %v\n", m)

		com := Command{m, c}
		comch <- com
	}
}

// GetCommand returns the command sent by the client
func GetCommand() (string, *net.TCPConn) {
	c := <-comch

	return c.Command, c.Conn
}

// Respond sends a response to the client
func Respond(c *net.TCPConn, r string) {
	c.Write([]byte(r))
}

// Kill stops the server
func Kill() {
	kill = true
}
