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
	r.HandleFunc("/", endpoints.RootHandler).Methods("GET")

	r.HandleFunc("/getpost", endpoints.GetPostHandler).Methods("GET")
	r.HandleFunc("/posts", endpoints.AllPostsHandler).Methods("GET")
	r.HandleFunc("/newpost", endpoints.NewPostHandler).Methods("POST")
	r.HandleFunc("/deletepost", endpoints.DeletePostHandler).Methods("POST")

	r.HandleFunc("/newuser", endpoints.NewUserHandler).Methods("POST")

	r.HandleFunc("/token", endpoints.TokenHandler).Methods("POST")
}
