package endpoints

import (
	"encoding/json"
	"goblog/internal/auth"
	"goblog/internal/models"
	"goblog/internal/utils"
	"net/http"
)

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	var possibleUser models.PossibleUser
	var R models.HttpResponse
	var code int

	err := utils.GetPostRequestData(r, &possibleUser)
	if err != nil {
		code = http.StatusUnprocessableEntity
		R.Status = "failure"
		R.Message = err.Error()
	} else {
		token, err := auth.GetToken(possibleUser)
		if err != nil {
			code = http.StatusUnauthorized
			R.Status = "failure"
			R.Message = err.Error()
		} else {
			token, err := json.Marshal(token)
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
	}
	utils.MakeResponse(w, code, R)
}

func TestTokenHandler(w http.ResponseWriter, r *http.Request) {
	var tokenPayload models.TokenPayload
	var R models.HttpResponse
	var code int
	valid := false

	err := utils.GetPostRequestData(r, &tokenPayload)
	if err != nil {
		code = http.StatusUnprocessableEntity
		R.Status = "failure"
		R.Message = err.Error()
	} else {
		if auth.CheckToken(tokenPayload.Username, tokenPayload.Token) {
			valid = true
		}
		valid, err := json.Marshal(models.TokenValidation{Valid: valid})
		if err != nil {
			code = http.StatusInternalServerError
			R.Status = "failure"
			R.Message = err.Error()
		} else {
			code = http.StatusOK
			R.Status = "success"
			R.Message = "token checked"
			R.Payload = json.RawMessage(valid)
		}
	}
	utils.MakeResponse(w, code, R)
}
