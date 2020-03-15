package endpoints

import (
	"encoding/json"
	"goblog/internal/database"
	"goblog/internal/utils"
	"net/http"
	"time"
)

type Post struct {
	Title         string `json:"title"`
	Posted        time.Time
	Posted_string string `json:"posted"`
	Body          string `json:"body"`
}

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	var post Post
	utils.GetPostRequestData(r, &post)
	post.Posted = time.Now()
	//todo ... DB stuff
}

func AllPostsHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	result, err := db.Query("SELECT * FROM public.post")
	if err != nil {
		panic(err)
	} else {
		defer result.Close()

		posts := []Post{}
		for result.Next() {
			var post Post
			err := result.Scan(&post.Title, &post.Posted_string, &post.Body)
			if err != nil {
				panic(err.Error())
			}
			posts = append(posts, post)
		}
		json.NewEncoder(w).Encode(posts)
	}

}
