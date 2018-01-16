// Package handler
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package handler

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/KaiserGald/rpgApp/services/logger"
	"github.com/KaiserGald/rpgApp/ui/handler/handle"
	"github.com/KaiserGald/rpgApp/ui/handler/handlers"
)

var routes []handle.Route

func init() {
	err := Add(index.Route())
	if err != nil {
		return
	}
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
func Handle(log *logger.Logger) {
	log.Debug.Log("Handling")
	for _, route := range routes {
		fmt.Println(route.Name())
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
