package endpoints

import (
	"goblog/internal/database"
	"goblog/internal/models"
	"goblog/internal/utils"
	"net/http"
)

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var R models.HttpResponse
	var code int

	err := utils.GetPostRequestData(r, &user)
	user.HashedPassword = utils.GenerateHash(user.HashedPassword)
	if err != nil {
		code = http.StatusUnprocessableEntity
		R.Status = "failure"
		R.Message = err.Error()
	} else {
		err = database.NewUser(user)
		if err != nil {
			code = http.StatusInternalServerError
			R.Status = "failure"
			R.Message = err.Error()
		} else {
			code = http.StatusCreated
			R.Status = "success"
			R.Message = "new user created"
		}
	}
	utils.MakeResponse(w, code, R)
}
