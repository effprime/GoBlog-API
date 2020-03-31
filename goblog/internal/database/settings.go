package database

import (
	"os"
	"strconv"
)

var host = os.Getenv("DB_ADDRESS")
var port, _ = strconv.Atoi(os.Getenv("DB_ACCESS_PORT"))
var user = os.Getenv("DB_USERNAME")
var password = os.Getenv("DB_PASSWORD")
var dbname = os.Getenv("DB_NAME")
