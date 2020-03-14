package main

import (
	"log"
	"net/http"
	"time"

	"goblog/internal/database"

	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	db := database.GetDatabaseConn()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
