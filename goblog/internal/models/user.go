package models

type User struct {
	Username       string `json:"username"`
	HashedPassword string `json:"password"`
	Email          string `json:"email"`
}

type PossibleUser struct {
	Username      string `json:"username"`
	PlainPassword string `json:"password"`
}
