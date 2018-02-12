// Package index
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package index

import (
	"net/http"

	"github.com/KaiserGald/unlicht-server/router/handler/handle"
)

// Route is the route that will be used
var route handle.Route

func init() {
	route = handle.Route{}
	route.SetName("/")
	route.SetHandler(handleFunc)
}

// handleFunc is the actual handler for the function
func handleFunc(w http.ResponseWriter, r *http.Request) {
	route.Log().Debug.Log("Handling Route '/'.\n")
	http.ServeFile(w, r, "/srv/unlichtServer/app/static/index.html")
}

// Route returns a pointer to the route
func Route() *handle.Route {
	return &route
}
