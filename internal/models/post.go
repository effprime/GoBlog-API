package models

type Post struct {
	Title  string `json:"title"`
	Posted string `json:"posted"`
	Body   string `json:"body"`
	Id     int    `json:"id"`
}
