package main

import (
	"goblog/internal/database"
	"goblog/internal/server"
)

func main() {
	database.TestDatabaseConn()
	server.Start()
}
