package server

import (
	"goblog/internal/router"
	"log"
	"net/http"
	"time"
)

func Start() int {
	address := "0.0.0.0:" + HTTP_PORT
	r := router.GetRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	return 0
}
