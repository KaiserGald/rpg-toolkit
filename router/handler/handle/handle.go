// Package handler
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package handle

import (
	"fmt"
	"net/http"

	"github.com/KaiserGald/logger"
)

// Route contains data about a specific route
type Route struct {
	name    string
	handler http.Handler
	log     *logger.Logger
}

// Name returns the name of the Route
func (r *Route) Name() string {
	return r.name
}

// SetName sets the name of the route
func (r *Route) SetName(s string) {
	r.name = s
}

// Handler returns the handler of the Route
func (r *Route) Handler() http.Handler {
	return r.handler
}

// SetHandler sets the handler function for the route
func (r *Route) SetHandler(fn http.HandlerFunc) {
	r.handler = fn
}

// Log returns a pointer to the logger
func (r *Route) Log() *logger.Logger {
	return r.log
}

// Init initializes the route
func (r *Route) Init(lg *logger.Logger) {
	r.log = lg
	r.log.Debug.Log("Route '%v' found!", r.Name())
}

func (r *Route) ErrorHandler(w http.ResponseWriter, req *http.Request, status int) {
	w.WriteHeader(status)
	fmt.Println(w, "404")
}
