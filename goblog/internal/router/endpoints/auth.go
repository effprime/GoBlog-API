package endpoints

import (
	"encoding/json"
	"goblog/internal/database"
	"goblog/internal/models"
	"goblog/internal/utils"
	"net/http"
)

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	var R models.HttpResponse
	var code int
	var possibleUser models.PossibleUser

	err := utils.GetPostRequestData(r, &possibleUser)
	if err != nil {
		code = http.StatusUnprocessableEntity
		R.Status = "failure"
		R.Message = err.Error()
	} else {
		token, err := database.GetToken(possibleUser)
		if err != nil {
			code = http.StatusInternalServerError
			R.Status = "failure"
			R.Message = err.Error()
		} else {
			code = http.StatusCreated
			R.Status = "success"
			R.Message = "user authenticated"
			R.Payload = json.RawMessage(token)
		}
	}
	utils.MakeResponse(w, code, R)
}
