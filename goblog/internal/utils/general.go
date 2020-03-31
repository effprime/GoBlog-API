package utils

import (
	"encoding/json"
	"goblog/internal/models"
	"net/http"
)

func GetPostRequestData(r *http.Request, class interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(class)
	if err != nil {
		return err
	}
	return nil
}

func MakeResponse(w http.ResponseWriter, code int, R models.HttpResponse) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(R)
}
