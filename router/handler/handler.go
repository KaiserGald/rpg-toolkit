// Package handler
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package handler

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/KaiserGald/logger"
	"github.com/KaiserGald/unlicht-server/router/handler/handle"
	"github.com/KaiserGald/unlicht-server/router/handler/handlers/index"
	"github.com/KaiserGald/unlicht-server/router/handler/handlers/testpage"
)

var (
	routes []handle.Route
	log    *logger.Logger
)

// Start starts the handler
func Start(lg *logger.Logger) error {
	log = lg
	log.Debug.Log("Starting route handler.")
	index.Route().Init(log)
	err := Add(index.Route())
	if err != nil {
		return err
	}

	testpage.Route().Init(log)
	err = Add(testpage.Route())
	if err != nil {
		return err
	}

	return nil
}

// Add adds a new route to the handler
func Add(r *handle.Route) error {
	log.Debug.Log("Adding '%v' to handler.", r.Name())
	if compareRoute(r) {
		return errors.New("Route already exists")
	}
	routes = append(routes, *r)
	return nil
}

// Handle handles all the registered routes
func Handle() {
	log.Debug.Log("Route Handler Started.")
	for _, route := range routes {
		http.Handle(route.Name(), route.Handler())
	}
}

// compareRoute checks to see if a given route already exists
func compareRoute(r *handle.Route) bool {
	for i := range routes {
		if routes[i].Name() == r.Name() {
			return true
		}
	}
	return false
}

func validRouteName(s string) (bool, error) {
	r, err := regexp.MatchString("^/[a-zA-Z]+\\w+$", s)
	if err != nil {
		return r, err
	}
	return r, nil
}
