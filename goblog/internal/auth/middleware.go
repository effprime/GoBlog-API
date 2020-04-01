package auth

import (
	"goblog/internal/models"
	"goblog/internal/utils"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			for _, path := range ALLOWED_PATHS {
				if path == r.URL.Path {
					next.ServeHTTP(w, r)
					return
				}
			}
			var token_payload models.TokenPayload
			utils.GetPostRequestData(r, &token_payload)
			if saved_token, ok := session_tokens[token_payload.Username]; ok {
				if validateHash(token_payload.Token.Value, saved_token.Value) {
					next.ServeHTTP(w, r)
				} else {
					http.Error(w, "API key does not match", http.StatusForbidden)
				}
			} else {
				http.Error(w, "Account not logged in", http.StatusForbidden)
			}
		})
}
