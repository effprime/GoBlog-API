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
	r.HandleFunc("/getpost", endpoints.GetPostHandler)
	r.HandleFunc("/posts", endpoints.AllPostsHandler)
	r.HandleFunc("/newpost", endpoints.NewPostHandler)
	r.HandleFunc("/deletepost", endpoints.DeletePostHandler)
}
