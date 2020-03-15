package server

import (
	"goblog/internal/router"
	"log"
	"net/http"
	"time"
)

func Start() int {
	r := router.GetRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	return 0
}
