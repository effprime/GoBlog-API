package router

import (
	"goblog/internal/router/endpoints"

	"github.com/gorilla/mux"
)

//GetRouter returns a final router object for API use
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	saturateRoutes(r)
	return r
}

//saturateRoutes adds routes to the router
func saturateRoutes(r *mux.Router) {
	r.HandleFunc("/", endpoints.RootHandler)
}
