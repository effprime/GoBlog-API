package database

import (
	"os"
	"strconv"
)

var host = os.Getenv("DB_HOST")
var port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
var user = os.Getenv("DB_USER")
var password = os.Getenv("DB_PASS")
var dbname = os.Getenv("DB_NAME")
