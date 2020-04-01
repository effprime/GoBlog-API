package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"goblog/internal/database"
	"goblog/internal/models"
)

const INCORRECT_ERROR string = "invalid username and/or password"

var session_tokens = make(map[string]models.Token)

func generateToken() models.Token {
	b := make([]byte, 8)
	rand.Read(b)
	token_string := fmt.Sprintf("%x", b)
	token := models.Token{Value: token_string}
	return token
}

func GetToken(possibleUser models.PossibleUser) (models.Token, error) {
	user, err := database.GetUser(possibleUser.Username)
	if err != nil {
		return models.Token{}, errors.New(INCORRECT_ERROR)
	}
	if validateHash(possibleUser.PlainPassword, user.HashedPassword) {
		token := generateToken()
		session_tokens[user.Username] = token
		return token, nil
	} else {
		return models.Token{}, errors.New(INCORRECT_ERROR)
	}
}
