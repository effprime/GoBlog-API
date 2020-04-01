package router

import "goblog/internal/utils"

var AUTH_REQUIRED = utils.GetEnv("AUTH_REQUIRED", "1")
