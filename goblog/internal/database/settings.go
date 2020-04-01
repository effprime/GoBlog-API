package database

import (
	"goblog/internal/utils"
	"strconv"
)

var DB_HOST = utils.GetEnv("DB_ADDRESS", "")
var DB_PORT, _ = strconv.Atoi(utils.GetEnv("DB_ACCESS_PORT", ""))
var DB_USER = utils.GetEnv("DB_USERNAME", "")
var DB_PASSWORD = utils.GetEnv("DB_PASSWORD", "")
var DB_NAME = utils.GetEnv("DB_NAME", "")
