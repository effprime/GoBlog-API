package database

import (
	"errors"
	"goblog/internal/models"
	"goblog/internal/utils"
)

const INCORRECT_MESSAGE string = "username or password is incorrect"

func GetToken(possibleUser models.PossibleUser) (string, error) {
	db := getDatabase()
	var user models.User

	err := db.QueryRow(
		"SELECT * FROM public.user where username=$1",
		possibleUser.Username,
	).Scan(&user.Username, &user.HashedPassword, &user.Email)
	if err != nil {
		return "", errors.New(INCORRECT_MESSAGE)
	}

	if utils.ValidateHash(possibleUser.PlainPassword, user.HashedPassword) {
		token := utils.GenerateToken()
		token_hash := utils.GenerateHash(token)
		_, err := db.Exec(
			"UPDATE public.user SET token_hash=$1 WHERE username=$2",
			token_hash,
			user.Username,
		)
		if err != nil {
			return token, nil
		} else {
			return "", err
		}
	} else {
		return "", errors.New(INCORRECT_MESSAGE)
	}
}
