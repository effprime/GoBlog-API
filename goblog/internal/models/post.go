package models

type Post struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Posted   string `json:"posted"`
	Body     string `json:"body"`
	Id       int    `json:"id"`
}
