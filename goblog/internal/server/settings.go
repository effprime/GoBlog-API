package server

import "goblog/internal/utils"

var HTTP_PORT = utils.GetEnv("CONTAINER_PORT", "8000")
