package models

type Token struct {
	Value string `json:"value"`
}

type TokenPayload struct {
	Token    Token  `json:"token"`
	Username string `json:"username"`
}
