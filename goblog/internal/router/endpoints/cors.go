package endpoints

import (
	"goblog/internal/models"
	"goblog/internal/utils"
	"net/http"
)

func HandleCorsRequest(w http.ResponseWriter, r *http.Request) {
	var R models.HttpResponse
	R.Status = "success"
	R.Message = ""
	utils.MakeResponse(w, 200, R)
}
