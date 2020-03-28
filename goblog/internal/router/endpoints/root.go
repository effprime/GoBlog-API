package endpoints

import "net/http"

//RootHandler handles requests to the API root endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
