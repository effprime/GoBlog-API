package utils

import (
	"encoding/json"
	"goblog/internal/models"
	"net/http"
	"os"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(R)
}

func GetEnv(lookup string, fallback string) string {
	value := os.Getenv(lookup)
	if value == "" {
		value = fallback
	}
	return value
}
