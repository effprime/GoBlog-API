package server

import (
	"goblog/internal/router"
	"net/http"
	"time"
)

//GetServer returns an API server interface
func GetServer() *http.Server {
	r := router.GetRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv
}
