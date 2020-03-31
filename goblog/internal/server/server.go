package server

import (
	"goblog/internal/router"
	"log"
	"net/http"
	"os"
	"time"
)

var port = os.Getenv("CONTAINER_PORT")

func Start() int {
	address := "0.0.0.0:" + port
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
