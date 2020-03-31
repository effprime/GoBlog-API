package database

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func getDatabase() *sql.DB {
	if db == nil {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
		var err error
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
	}
	return db
}

func TestDatabaseConn() {
	db = getDatabase()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
}
