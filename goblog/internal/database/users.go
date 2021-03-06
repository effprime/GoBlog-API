package database

import (
	"errors"
	"goblog/internal/models"
)

func NewUser(user models.User) error {
	_, err := GetUser(user.Username)
	if err == nil {
		return errors.New("user already exists")
	}
	db := getDatabase()
	_, err = db.Exec("INSERT INTO blog.person (username, hash_password, email) VALUES ($1, $2, $3)", user.Username, user.HashedPassword, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(username string) (models.User, error) {
	var user models.User
	db := getDatabase()
	err := db.QueryRow(
		"SELECT * FROM blog.person where username=$1",
		username,
	).Scan(&user.Username, &user.HashedPassword, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}
