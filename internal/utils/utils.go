package utils

import (
	"encoding/json"
	"net/http"
)

func GetPostRequestData(r *http.Request, class interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(class)
	if err != nil {
		panic(err)
	}
}
