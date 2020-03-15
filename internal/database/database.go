package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "goblog"
)

var db *sql.DB

func GetDatabase() *sql.DB {
	if db == nil {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		return db
	}
	return db
}

func TestDatabaseConn() int {
	db = GetDatabase()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	return 0
}
