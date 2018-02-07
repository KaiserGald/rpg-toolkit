// Package handler
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package handler

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/KaiserGald/unlichtServer/services/logger"
	"github.com/KaiserGald/unlichtServer/ui/handler/handle"
	"github.com/KaiserGald/unlichtServer/ui/handler/handlers"
)

var (
	routes []handle.Route
	log    *logger.Logger
)

// Start starts the handler
func Start(lg *logger.Logger) error {
	log = lg
	index.Route().Init(log)
	err := Add(index.Route())
	if err != nil {
		return err
	}

	return nil
}

// Add adds a new route to the handler
func Add(r *handle.Route) error {
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
