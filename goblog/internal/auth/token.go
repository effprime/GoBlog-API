package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"goblog/internal/database"
	"goblog/internal/models"
)

var session_tokens = make(map[string]models.Token)

func generateToken() models.Token {
	b := make([]byte, 8)
	rand.Read(b)
	return models.Token{Value: fmt.Sprintf("%x", b)}
}

func GetToken(possibleUser models.PossibleUser) (models.Token, error) {
	incorrect_error := "invalid username and/or password"

	user, err := database.GetUser(possibleUser.Username)
	if err != nil {
		return models.Token{}, errors.New(incorrect_error)
	}
	if validateHash(possibleUser.PlainPassword, user.HashedPassword) {
		token := generateToken()
		session_tokens[user.Username] = models.Token{Value: GenerateHash(token.Value)}
		return token, nil

	} else {
		return models.Token{}, errors.New(incorrect_error)
	}
}

func CheckToken(username string, token models.Token) bool {
	if savedToken, ok := session_tokens[username]; ok {
		if validateHash(token.Value, savedToken.Value) {
			return true
		}
	}
	return false
}
