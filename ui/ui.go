package ui

import (
	"net"
	"net/http"
	"time"

	"github.com/KaiserGald/rpgApp/services/logger"
)

// Start initializes and starts the ui router
func Start(listener net.Listener, log *logger.Logger) {

	log.Info.Log("Starting front-end.\n")

	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	go server.Serve(listener)

	log.Info.Log("Server up and running.\n")
}
