package utils

import (
	"encoding/json"
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

func VerifyMethod(r *http.Request, method_target string) bool {
	if r.Method == method_target {
		return true
	}
	return false
}
