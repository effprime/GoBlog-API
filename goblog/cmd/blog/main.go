package main

import (
	"goblog/internal/database"
	"goblog/internal/server"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	log.Output(0, "Checking database connection...")
	database.TestDatabaseConn()

	log.Output(0, "Starting server...")
	server.Start()
}
