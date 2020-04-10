package router

import "goblog/internal/utils"

var AUTH_REQUIRED = utils.GetEnv("AUTH_REQUIRED", "1")
var ALLOW_CORS = utils.GetEnv("ALLOW_CORS", "1")
