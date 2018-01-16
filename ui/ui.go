// Package ui
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package ui

import (
	"net"
	"net/http"
	"time"

	"github.com/KaiserGald/rpgApp/services/logger"
	"github.com/KaiserGald/rpgApp/ui/handler"
)

// Start initializes and starts the ui router
func Start(listener net.Listener, log *logger.Logger) {

	log.Info.Log("Starting front-end.\n")

	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	handler.Handle(log)
	go server.Serve(listener)

	log.Info.Log("Front-end up and running.\n")
}
