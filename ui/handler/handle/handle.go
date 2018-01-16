// Package handler
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package handle

import (
	"net/http"
)

// Route contains data about a specific route
type Route struct {
	name    string
	handler http.Handler
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
