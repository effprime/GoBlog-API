package main

import (
	"goblog/internal/database"
	"goblog/internal/server"
	"log"
)

func main() {
	db := database.GetDatabaseConn()
	defer db.Close()

	srv := server.GetServer()
	log.Fatal(srv.ListenAndServe())
}
