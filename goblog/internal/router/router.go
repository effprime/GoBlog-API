package router

import (
	"goblog/internal/auth"
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
	r.HandleFunc("/auth/token", endpoints.TokenHandler).Methods("POST")
	r.HandleFunc("/api/getpost", endpoints.GetPostHandler).Methods("GET")
	r.HandleFunc("/api/posts", endpoints.AllPostsHandler).Methods("GET")
	r.HandleFunc("/api/newpost", endpoints.NewPostHandler).Methods("POST")
	r.HandleFunc("/api/deletepost", endpoints.DeletePostHandler).Methods("POST")
	r.HandleFunc("/api/newuser", endpoints.NewUserHandler).Methods("POST")

	if AUTH_REQUIRED == "1" {
		r.Use(auth.Middleware)
	}
}
