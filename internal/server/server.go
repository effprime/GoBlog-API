package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//Handles requests to the API root endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//GetServer returns an API server interface
func GetServer() *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv
}
