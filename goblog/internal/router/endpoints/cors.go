package endpoints

import "net/http"

func HandleCorsRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, content-type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,DELETE,PUT,OPTIONS")
}
