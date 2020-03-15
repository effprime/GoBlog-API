package router

import (
	"goblog/internal/router/endpoints"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	saturateRoutes(r)
	return r
}

func saturateRoutes(r *mux.Router) {
	r.HandleFunc("/", endpoints.RootHandler)
	r.HandleFunc("/posts", endpoints.AllPostsHandler)
	r.HandleFunc("/newpost", endpoints.NewPostHandler)
}
